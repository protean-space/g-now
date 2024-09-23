package cron

import (
	"context"
	"encoding/xml"
	"fmt"
	"g-now/db"
	"g-now/model"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func FetchNews() {
	var categoryMaster = []string{"政治", "ビジネス", "テクノロジー", "エンタメ", "スポーツ", "天気"}
	var err error

	// URLからXMLデータを取得
	url := "https://news.google.com/rss?hl=ja&gl=JP&ceid=JP:ja"
	resp, err := http.Get(url)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	// XMLをパースする
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// .envをロード
	err = godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	// Geminiを初期化
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		slog.Error("error: connect Gemini.")
		slog.Error(err.Error())
	}
	defer client.Close()

	// DB初期化
	dbConn, err := db.NewDB()
	if err != nil {
		slog.Error("fail connect db.")
		panic(err)
	}
	defer db.CloseDB(dbConn, err)

	genModel := client.GenerativeModel("gemini-1.5-flash")
	categorizePrompt := "以下のニュースタイトルを、次のcategory: の中から1つに分類分けしてください。出力は結果の1単語だけを表示してください。category: [%s]。\nニュースタイトル: %s"
	summarizePrompt := "以下のXMLが述べていることを300字以内で要約して出力してください。出力は要約内容だけを表示してください。\nXML: %s"
	taggingPrompt := "以下のニュースタイトルから、カテゴリを3つ以内で作成してください。カテゴリは,区切りの結果だけを出力してください。\nニュースタイトル: %s"

	slog.Info("start fetch google news...")
	slog.Info(os.Getenv("GEMINI_API_KEY"))

	for index, item := range rss.Channel.Items {
		// 5件まで作成
		if index == 5 {
			return
		}

		var existsRecord []model.Article
		dbConn.Where("uid = ?", item.GUID).Find(&existsRecord)
		if len(existsRecord) != 0 {
			// すでに登録されている記事ならスキップ
			slog.Info("skip article UID: " + item.GUID)
			continue
		}

		// カテゴリ
		var category model.Category
		categoryRes, err := genModel.GenerateContent(ctx, genai.Text(fmt.Sprintf(categorizePrompt, strings.Join(categoryMaster, ","), item.Title)))
		if err != nil {
			slog.Error(err.Error())
		}
		if categoryRes != nil {
			generateCategory := categoryRes.Candidates[0].Content.Parts[0]
			dbConn.Where("category_name = ?", strings.TrimSpace(fmt.Sprint(generateCategory))).First(&category)
		}

		// コンテンツ
		summaryRes, err := genModel.GenerateContent(ctx, genai.Text(fmt.Sprintf(summarizePrompt, item.Description)))
		if err != nil {
			slog.Error(err.Error())
		}
		generateSummary := summaryRes.Candidates[0].Content.Parts[0]
		summary := strings.TrimSpace(fmt.Sprint(generateSummary))

		// タグ
		var tagIds []uint
		tagRes, err := genModel.GenerateContent(ctx, genai.Text(fmt.Sprintf(taggingPrompt, item.Title)))
		if err != nil {
			slog.Error(err.Error())
		}
		generateTags := tagRes.Candidates[0].Content.Parts[0]
		for _, s := range strings.Split(strings.TrimSpace(fmt.Sprint(generateTags)), ",") {
			var tag model.Tag
			dbConn.Where(model.Tag{TagName: s}).FirstOrCreate(&tag, model.Tag{TagName: s})
			tagIds = append(tagIds, tag.ID)
		}

		// 記事の投稿時間
		parsedTime, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			slog.Error(err.Error())
		}

		// 記事の作成
		article := model.Article{UID: item.GUID, Title: item.Title, Contents: summary, ArticleUrl: item.Link, SourcePublishedAt: parsedTime}
		dbConn.Create(&article)

		// 記事とタグの中間テーブルにレコードを作成
		for _, tagId := range tagIds {
			articleTagMap := model.ArticleTagMap{ArticleId: article.ID, TagId: tagId}
			dbConn.Create(&articleTagMap)
		}
		// 記事とカテゴリの中間テーブルにレコードを作成
		articleCategoryMap := model.ArticleCategoryMap{ArticleId: article.ID, CategoryId: category.ID}
		dbConn.Create(&articleCategoryMap)

		slog.Info("Successfully created article. UID: " + item.GUID)
	}
	slog.Info("end fetch google news.")
}
