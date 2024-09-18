# g-now

## 環境構築

### PostgreSQLコンテナの起動

```sh
cd backend
docker compose up
```

### マイグレーション実行

```sh
cd backend
GO_ENV=dev go run migrate/migrate.go
```

### Seed実行

```sh
GO_ENV=dev go run db/seed/seed.go
```

### GeminiAPIの設定

- [APIキーの作成](https://ai.google.dev/gemini-api/docs/get-started/tutorial?lang=go&authuser=3&%3Bhl=ja&hl=ja#:~:text=%E3%82%82%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%99%E3%80%82-,API%20%E3%82%AD%E3%83%BC%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B,-API%20%E3%82%AD%E3%83%BC%E3%82%92)
- .envファイルの作成
  `cp .env.sample .env`
- `.env`ファイルの`GEMINI_API_KEY`にAPIキーを記述

### cronの手動実行

```sh
GO_ENV=dev go run cron/main.go
```

### サーバー起動

```sh
cd backend
GO_ENV=dev go run main.go
```

http://localhost:8080/articles

<details>

<summary>Sample Response</summary>

`curl -s "localhost:8080/articles"`

```json
[
  {
    "id": 1,
    "uid": "aaa",
    "title": "「斎藤家の食卓をにぎわしただけ」〝おねだり〟疑惑に百条委員が指摘 兵庫知事は正当性主張 - 産経ニュース",
    "contents": "兵庫県の斎藤元彦知事が、県の事業で業者に便宜を図った見返りに飲食接待を受けたとされる疑惑で、批判が高まっています。斎藤知事は百条委員会で正当性を主張したものの、東国原英夫氏などから「知事としての資質はない」と厳しい声が上がっています。自民党を含む県議会3会派は辞職を求める方向で、維新以外の全会派が同調する見通しです。",
    "ArticleUrl": "https://news.google.com/rss/articles/CBMidkFVX3lxTFBKQ3huMkR2UTAxbVFSNmRaeG1ZTmpXbkFVWmhwdkZ6bTNSMWpSYUludWphWmZRZnVUY0U2d28wQ3FhWG5mcXlYblpIVV9FREhLbDkybFhtN0pNcy0xZ0k4ZmxhVXRpQW9tdXRkRXFfeDQ2UUhqYXc?oc=5",
    "ArticleImageUrl": "",
    "PageView": 0,
    "PayloadJson": "",
    "source_published_at": "2024-09-07T08:01:00+09:00",
    "CreatedAt": "2024-09-16T11:00:56.909902+09:00",
    "UpdatedAt": "2024-09-16T11:00:56.909902+09:00",
    "categories": [
      {
        "id": 1,
        "category_name": "政治"
      }
    ],
    "tags": [
      {
        "id": 1,
        "tag_name": "政治"
      },
      {
        "id": 2,
        "tag_name": "地方行政"
      },
      {
        "id": 3,
        "tag_name": "スキャンダル"
      }
    ]
  },
  {
    "id": 2,
    "uid": "bbb",
    "title": "「1年間で改革を」小泉進次郎氏が演説 野田聖子氏、立候補を模索 [自民] - 朝日新聞デジタル",
    "contents": "自民党総裁選への立候補を模索する野田聖子氏を支援するため、小泉進次郎氏が銀座で街頭演説を行いました。「1年間で改革を」と訴え、党改革への意欲を示しました。質疑応答では、フリー記者からの辛辣な質問に対し、冷静に切り返す場面も見られました。一方、その質問内容については「知的レベルが低い」と脳科学者から指摘されるなど、波紋も広がっています。",
    "ArticleUrl": "https://news.google.com/rss/articles/CBMiZ0FVX3lxTE52cGV5c2xaN003VnZnb1NlSUV2UkQ1dXRkME43RkRGME5uOENTcGtTdm5OeDVITlpkN1IxdFRpTlVWWTJYSDlaaklKY2hlM2RXazh1SWFPeVFGYnY4a3l6NFRMZDhPWnM?oc=5",
    "ArticleImageUrl": "",
    "PageView": 0,
    "PayloadJson": "",
    "source_published_at": "2024-09-07T08:45:00+09:00",
    "CreatedAt": "2024-09-16T11:00:56.911127+09:00",
    "UpdatedAt": "2024-09-16T11:00:56.911127+09:00",
    "categories": [
      {
        "id": 1,
        "category_name": "政治"
      }
    ],
    "tags": [
      {
        "id": 1,
        "tag_name": "政治"
      },
      {
        "id": 4,
        "tag_name": "自民党"
      },
      {
        "id": 5,
        "tag_name": "選挙"
      }
    ]
  },
  {
    "id": 3,
    "uid": "ccc",
    "title": "栃木 真岡市の音楽イベントで9人けが 会場周辺で落雷情報も - nhk.or.jp",
    "contents": "栃木県真岡市の井頭公園で開かれていた野外音楽イベントで、落雷とみられる事故が発生し、１０代から２０代の男女９人が手足のしびれを訴えるなどして負傷しました。会場周辺では当時、激しい雷雨に見舞われており、落雷の影響とみられています。イベントは中止となりました。",
    "ArticleUrl": "https://news.google.com/rss/articles/CBMib0FVX3lxTFBSQmxhb0trb09BRFYtLWxOSTVTVlEycjBKY0FQTk9BNnM5dTJsSG0yRVo5UXJzQ3VCdTl1TnVja0J4cm85d0t5TzA0RW1sQjFvdE9JYlhja2J6YWctR3hKY3V2LXA5S3BnYXJicExDaw?oc=5",
    "ArticleImageUrl": "",
    "PageView": 0,
    "PayloadJson": "",
    "source_published_at": "2024-09-07T19:45:00+09:00",
    "CreatedAt": "2024-09-16T11:00:56.911921+09:00",
    "UpdatedAt": "2024-09-16T11:00:56.911921+09:00",
    "categories": [
      {
        "id": 6,
        "category_name": "天気"
      }
    ],
    "tags": [
      {
        "id": 6,
        "tag_name": "事故"
      },
      {
        "id": 7,
        "tag_name": "気象"
      },
      {
        "id": 8,
        "tag_name": "イベント"
      }
    ]
  }
]
```

</details>

### 参考

#### DBログイン

```sh
psql -U user -d g-now
```

#### ER図

拡張機能：https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio

[ER図](er.drawio)
