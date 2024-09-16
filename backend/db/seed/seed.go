package main

import (
	"g-now/db"
	"g-now/model"
	"log/slog"
	"time"

	"gorm.io/gorm/clause"
)

func main() {
	slog.Info("Start seeding.")

	var err error
	dbConn, err := db.NewDB()
	if err != nil {
		slog.Error("fail connect db.")
		panic(err)
	}
	defer db.CloseDB(dbConn, err)

	dbConn.Migrator().DropTable(&model.ArticleTagMap{}, &model.Tag{}, &model.ArticleCategoryMap{}, &model.Category{}, &model.Article{})
	dbConn.AutoMigrate(&model.Article{}, &model.Category{}, &model.ArticleCategoryMap{}, &model.Tag{}, &model.ArticleTagMap{})

	categories := []model.Category{
		{ID: 1, CategoryName: "政治"},
		{ID: 2, CategoryName: "ビジネス"},
		{ID: 3, CategoryName: "テクノロジー"},
		{ID: 4, CategoryName: "エンタメ"},
		{ID: 5, CategoryName: "スポーツ"},
		{ID: 6, CategoryName: "天気"},
	}
	for _, category := range categories {
		err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&category).Error
		if err != nil {
			slog.Error("fail create categories data.")
		} else {
			slog.Info("success create categories.")
		}
	}

	tags := []model.Tag{
		{ID: 1, TagName: "政治"},
		{ID: 2, TagName: "地方行政"},
		{ID: 3, TagName: "スキャンダル"},
		{ID: 4, TagName: "自民党"},
		{ID: 5, TagName: "選挙"},
		{ID: 6, TagName: "事故"},
		{ID: 7, TagName: "気象"},
		{ID: 8, TagName: "イベント"},
	}
	for _, tag := range tags {
		err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&tag).Error
		if err != nil {
			slog.Error("fail create tags data.")
		} else {
			slog.Info("success create tags.")
		}
	}

	articles := []model.Article{
		{ID: 1, UID: "aaa", Title: "「斎藤家の食卓をにぎわしただけ」〝おねだり〟疑惑に百条委員が指摘 兵庫知事は正当性主張 - 産経ニュース", Contents: "兵庫県の斎藤元彦知事が、県の事業で業者に便宜を図った見返りに飲食接待を受けたとされる疑惑で、批判が高まっています。斎藤知事は百条委員会で正当性を主張したものの、東国原英夫氏などから「知事としての資質はない」と厳しい声が上がっています。自民党を含む県議会3会派は辞職を求める方向で、維新以外の全会派が同調する見通しです。", ArticleUrl: "https://news.google.com/rss/articles/CBMidkFVX3lxTFBKQ3huMkR2UTAxbVFSNmRaeG1ZTmpXbkFVWmhwdkZ6bTNSMWpSYUludWphWmZRZnVUY0U2d28wQ3FhWG5mcXlYblpIVV9FREhLbDkybFhtN0pNcy0xZ0k4ZmxhVXRpQW9tdXRkRXFfeDQ2UUhqYXc?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 8, 1, 0, 0, time.Local)},
		{ID: 2, UID: "bbb", Title: "「1年間で改革を」小泉進次郎氏が演説 野田聖子氏、立候補を模索 [自民] - 朝日新聞デジタル", Contents: "自民党総裁選への立候補を模索する野田聖子氏を支援するため、小泉進次郎氏が銀座で街頭演説を行いました。「1年間で改革を」と訴え、党改革への意欲を示しました。質疑応答では、フリー記者からの辛辣な質問に対し、冷静に切り返す場面も見られました。一方、その質問内容については「知的レベルが低い」と脳科学者から指摘されるなど、波紋も広がっています。", ArticleUrl: "https://news.google.com/rss/articles/CBMiZ0FVX3lxTE52cGV5c2xaN003VnZnb1NlSUV2UkQ1dXRkME43RkRGME5uOENTcGtTdm5OeDVITlpkN1IxdFRpTlVWWTJYSDlaaklKY2hlM2RXazh1SWFPeVFGYnY4a3l6NFRMZDhPWnM?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 8, 45, 0, 0, time.Local)},
		{ID: 3, UID: "ccc", Title: "栃木 真岡市の音楽イベントで9人けが 会場周辺で落雷情報も - nhk.or.jp", Contents: "栃木県真岡市の井頭公園で開かれていた野外音楽イベントで、落雷とみられる事故が発生し、１０代から２０代の男女９人が手足のしびれを訴えるなどして負傷しました。会場周辺では当時、激しい雷雨に見舞われており、落雷の影響とみられています。イベントは中止となりました。", ArticleUrl: "https://news.google.com/rss/articles/CBMib0FVX3lxTFBSQmxhb0trb09BRFYtLWxOSTVTVlEycjBKY0FQTk9BNnM5dTJsSG0yRVo5UXJzQ3VCdTl1TnVja0J4cm85d0t5TzA0RW1sQjFvdE9JYlhja2J6YWctR3hKY3V2LXA5S3BnYXJicExDaw?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 19, 45, 0, 0, time.Local)},
	}
	for _, article := range articles {
		err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&article).Error
		if err != nil {
			slog.Error("fail create articles data.")
		} else {
			slog.Info("success create articles.")
		}
	}

	articleCategoryMap := []model.ArticleCategoryMap{
		{ArticleId: articles[0].ID, CategoryId: categories[0].ID}, // カテゴリ：政治
		{ArticleId: articles[1].ID, CategoryId: categories[0].ID}, // カテゴリ：政治
		{ArticleId: articles[2].ID, CategoryId: categories[5].ID}, // カテゴリ：天気
	}
	if err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&articleCategoryMap).Error; err != nil {
		slog.Error("fail create article_category_maps data.")
	} else {
		slog.Info("success create article_category_maps.")
	}

	articleTagMap := []model.ArticleTagMap{
		{ArticleId: articles[0].ID, TagId: tags[0].ID},
		{ArticleId: articles[0].ID, TagId: tags[1].ID},
		{ArticleId: articles[0].ID, TagId: tags[2].ID},
		{ArticleId: articles[1].ID, TagId: tags[0].ID},
		{ArticleId: articles[1].ID, TagId: tags[3].ID},
		{ArticleId: articles[1].ID, TagId: tags[4].ID},
		{ArticleId: articles[2].ID, TagId: tags[5].ID},
		{ArticleId: articles[2].ID, TagId: tags[6].ID},
		{ArticleId: articles[2].ID, TagId: tags[7].ID},
	}
	for _, article_tag := range articleTagMap {
		err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&article_tag).Error
		if err != nil {
			slog.Error("fail create articles data.")
		} else {
			slog.Info("success create articles.")
		}
	}

	slog.Info("Finish seeding.")
}
