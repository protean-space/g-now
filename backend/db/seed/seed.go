package seed

import (
	"g-now/db"
	"g-now/model"
	"log/slog"
	"time"

	"gorm.io/gorm/clause"
)

func Run() {
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
		{CategoryName: "政治"},
		{CategoryName: "エンタメ"},
		{CategoryName: "経済"},
		{CategoryName: "スポーツ"},
		{CategoryName: "ビジネス"},
		{CategoryName: "天気"},
		{CategoryName: "キャリア"},
		{CategoryName: "テクノロジー"},
		{CategoryName: "ゲーム"},
		{CategoryName: "グルメ"},
		{CategoryName: "Web3"},
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
		{TagName: "政治"},
		{TagName: "地方行政"},
		{TagName: "スキャンダル"},
		{TagName: "自民党"},
		{TagName: "選挙"},
		{TagName: "事故"},
		{TagName: "気象"},
		{TagName: "イベント"},
		{TagName: "芸能人"},
		{TagName: "大谷翔平"},
		{TagName: "不倫"},
		{TagName: "告白"},
		{TagName: "晴れ"},
		{TagName: "人身事故"},
		{TagName: "電車遅延"},
		{TagName: "民主党"},
		{TagName: "社会人"},
		{TagName: "シゴデキ"},
		{TagName: "ギター"},
		{TagName: "2040年"},
		{TagName: "少子化"},
		{TagName: "労働"},
		{TagName: "外交"},
		{TagName: "出張"},
		{TagName: "キャバクラ"},
		{TagName: "スマホ"},
		{TagName: "iPhone"},
		{TagName: "POP-UP ショップ"},
		{TagName: "コスメ"},
		{TagName: "美容"},
		{TagName: "うなぎ"},
		{TagName: "飲食店"},
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
		{UID: "aaa", Title: "「斎藤家の食卓をにぎわしただけ」〝おねだり〟疑惑に百条委員が指摘 兵庫知事は正当性主張 - 産経ニュース", Contents: "兵庫県の斎藤元彦知事が、県の事業で業者に便宜を図った見返りに飲食接待を受けたとされる疑惑で、批判が高まっています。斎藤知事は百条委員会で正当性を主張したものの、東国原英夫氏などから「知事としての資質はない」と厳しい声が上がっています。自民党を含む県議会3会派は辞職を求める方向で、維新以外の全会派が同調する見通しです。", ArticleUrl: "https://news.google.com/rss/articles/CBMidkFVX3lxTFBKQ3huMkR2UTAxbVFSNmRaeG1ZTmpXbkFVWmhwdkZ6bTNSMWpSYUludWphWmZRZnVUY0U2d28wQ3FhWG5mcXlYblpIVV9FREhLbDkybFhtN0pNcy0xZ0k4ZmxhVXRpQW9tdXRkRXFfeDQ2UUhqYXc?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 8, 1, 0, 0, time.Local)},
		{UID: "bbb", Title: "「1年間で改革を」小泉進次郎氏が演説 野田聖子氏、立候補を模索 [自民] - 朝日新聞デジタル", Contents: "自民党総裁選への立候補を模索する野田聖子氏を支援するため、小泉進次郎氏が銀座で街頭演説を行いました。「1年間で改革を」と訴え、党改革への意欲を示しました。質疑応答では、フリー記者からの辛辣な質問に対し、冷静に切り返す場面も見られました。一方、その質問内容については「知的レベルが低い」と脳科学者から指摘されるなど、波紋も広がっています。", ArticleUrl: "https://news.google.com/rss/articles/CBMiZ0FVX3lxTE52cGV5c2xaN003VnZnb1NlSUV2UkQ1dXRkME43RkRGME5uOENTcGtTdm5OeDVITlpkN1IxdFRpTlVWWTJYSDlaaklKY2hlM2RXazh1SWFPeVFGYnY4a3l6NFRMZDhPWnM?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 8, 45, 0, 0, time.Local)},
		{UID: "ccc", Title: "栃木 真岡市の音楽イベントで9人けが 会場周辺で落雷情報も - nhk.or.jp", Contents: "栃木県真岡市の井頭公園で開かれていた野外音楽イベントで、落雷とみられる事故が発生し、１０代から２０代の男女９人が手足のしびれを訴えるなどして負傷しました。会場周辺では当時、激しい雷雨に見舞われており、落雷の影響とみられています。イベントは中止となりました。", ArticleUrl: "https://news.google.com/rss/articles/CBMib0FVX3lxTFBSQmxhb0trb09BRFYtLWxOSTVTVlEycjBKY0FQTk9BNnM5dTJsSG0yRVo5UXJzQ3VCdTl1TnVja0J4cm85d0t5TzA0RW1sQjFvdE9JYlhja2J6YWctR3hKY3V2LXA5S3BnYXJicExDaw?oc=5", SourcePublishedAt: time.Date(2024, 9, 7, 19, 45, 0, 0, time.Local)},
		{UID: "ddd", Title: "「もうテレビで見たくない」不倫男性タレントは？3位は小室哲哉、2位は渡部建、圧倒的1位は「人の道として許せない」【500人に聞く】", Contents: "　一度世間に知られてしまえば、それまで築き上げたイメージは崩壊し、業界での立ち回りも変化せざるを得ない芸能人の不倫報道。【第5位】KAT-TUN・中丸雄一：38票　【第4位】ダウンタウン・松本人志：45票 【第3位】小室哲哉：49票 【第2位】アンジャッシュ・渡部建：91票 【第1位】東出昌大：154票　元妻・杏が第三子を妊娠・出産するなか、女優・唐田えりかと2年半にわたり不倫していたことが報じられた。　軽はずみな行動の代償は計り知れない。", ArticleUrl: "https://news.yahoo.co.jp/articles/10c3d9986f7fe2a2ea4beb6343ad27885e50fa61", SourcePublishedAt: time.Date(2024, 9, 21, 14, 56, 0, 0, time.Local)},
		{UID: "eee", Title: "”2040年に目指すべき日本” 経団連 ビジョン年内とりまとめへ", Contents: "経団連の夏のフォーラムが18日から始まりました。経団連の十倉会長は、人口減少や少子高齢化が進む現状などを踏まえながら議論を深め、2040年に目指すべき日本の経済や社会のあり方について、年内にビジョンをとりまとめる考えを示しました。", ArticleUrl: "https://www3.nhk.or.jp/news/html/20240718/k10014515001000.html", SourcePublishedAt: time.Date(2024, 7, 18, 21, 12, 0, 0, time.Local)},
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
		{ArticleId: 1, CategoryId: 1}, // カテゴリ：政治
		{ArticleId: 2, CategoryId: 1}, // カテゴリ：政治
		{ArticleId: 3, CategoryId: 6}, // カテゴリ：天気
		{ArticleId: 4, CategoryId: 2}, // カテゴリ：エンタメ
		{ArticleId: 5, CategoryId: 5}, // カテゴリ：ビジネル
	}
	if err := dbConn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&articleCategoryMap).Error; err != nil {
		slog.Error("fail create article_category_maps data.")
	} else {
		slog.Info("success create article_category_maps.")
	}

	articleTagMap := []model.ArticleTagMap{
		{ArticleId: 1, TagId: 1},
		{ArticleId: 1, TagId: 2},
		{ArticleId: 1, TagId: 3},
		{ArticleId: 2, TagId: 1},
		{ArticleId: 2, TagId: 4},
		{ArticleId: 2, TagId: 5},
		{ArticleId: 3, TagId: 6},
		{ArticleId: 3, TagId: 7},
		{ArticleId: 3, TagId: 8},
		{ArticleId: 4, TagId: 3},
		{ArticleId: 4, TagId: 25},
		{ArticleId: 4, TagId: 9},
		{ArticleId: 5, TagId: 20},
		{ArticleId: 5, TagId: 21},
		{ArticleId: 5, TagId: 22},
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
