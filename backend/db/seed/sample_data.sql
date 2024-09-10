INSERT INTO categories (`id`, `category_name`) VALUES (1, '政治'), (2, 'ビジネス'), (3, 'テクノロジー'), (4, 'エンタメ'), (5, 'スポーツ'), (6, '天気');

INSERT INTO articles (`id`, `uid`, `title`, `contents`, `article_url`, `source_published_at`) VALUES
(1, 'aaaaa', '「斎藤家の食卓をにぎわしただけ」〝おねだり〟疑惑に百条委員が指摘 兵庫知事は正当性主張 - 産経ニュース', '兵庫県の斎藤元彦知事が、県の事業で業者に便宜を図った見返りに飲食接待を受けたとされる疑惑で、批判が高まっています。斎藤知事は百条委員会で正当性を主張したものの、東国原英夫氏などから「知事としての資質はない」と厳しい声が上がっています。自民党を含む県議会3会派は辞職を求める方向で、維新以外の全会派が同調する見通しです。', 'https://news.google.com/rss/articles/CBMidkFVX3lxTFBKQ3huMkR2UTAxbVFSNmRaeG1ZTmpXbkFVWmhwdkZ6bTNSMWpSYUludWphWmZRZnVUY0U2d28wQ3FhWG5mcXlYblpIVV9FREhLbDkybFhtN0pNcy0xZ0k4ZmxhVXRpQW9tdXRkRXFfeDQ2UUhqYXc?oc=5', '2024-09-07 08:01:00'),
(2, 'bbbbb', '「1年間で改革を」小泉進次郎氏が演説 野田聖子氏、立候補を模索 [自民] - 朝日新聞デジタル', '自民党総裁選への立候補を模索する野田聖子氏を支援するため、小泉進次郎氏が銀座で街頭演説を行いました。「1年間で改革を」と訴え、党改革への意欲を示しました。質疑応答では、フリー記者からの辛辣な質問に対し、冷静に切り返す場面も見られました。一方、その質問内容については「知的レベルが低い」と脳科学者から指摘されるなど、波紋も広がっています。', 'https://news.google.com/rss/articles/CBMiZ0FVX3lxTE52cGV5c2xaN003VnZnb1NlSUV2UkQ1dXRkME43RkRGME5uOENTcGtTdm5OeDVITlpkN1IxdFRpTlVWWTJYSDlaaklKY2hlM2RXazh1SWFPeVFGYnY4a3l6NFRMZDhPWnM?oc=5', '2024-09-07 08:45:00'),
(3, 'ccccc', '栃木 真岡市の音楽イベントで9人けが 会場周辺で落雷情報も - nhk.or.jp', '栃木県真岡市の井頭公園で開かれていた野外音楽イベントで、落雷とみられる事故が発生し、１０代から２０代の男女９人が手足のしびれを訴えるなどして負傷しました。会場周辺では当時、激しい雷雨に見舞われており、落雷の影響とみられています。イベントは中止となりました。', 'https://news.google.com/rss/articles/CBMib0FVX3lxTFBSQmxhb0trb09BRFYtLWxOSTVTVlEycjBKY0FQTk9BNnM5dTJsSG0yRVo5UXJzQ3VCdTl1TnVja0J4cm85d0t5TzA0RW1sQjFvdE9JYlhja2J6YWctR3hKY3V2LXA5S3BnYXJicExDaw?oc=5', '2024-09-07 19:45:00');

INSERT INTO article_category_maps (`article_id`, `category_id`), VALUES (1, 1), (2, 1), (3, 6);

-- tagsとarticlesの中間テーブルで作成するか検討
INSERT INTO tags (`article_id`, `tag_name`) VALUES (1, '政治'), (1, '地方行政'), (1, 'スキャンダル'), (2, '政治'), (2, '自民党'), (2, '選挙'), (3, '事故'), (3, '気象'), (3, 'イベント');
