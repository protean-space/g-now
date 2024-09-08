package cron

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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

func main() {
	// URLからXMLデータを取得
	url := "https://news.google.com/rss?hl=ja&gl=JP&ceid=JP:ja"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("XMLデータの取得に失敗しました:", err)
		return
	}
	defer resp.Body.Close()

	// XMLをパースする
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("XMLのパースに失敗しました:", err)
		return
	}

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		fmt.Println("XMLのアンマーシャルに失敗しました:", err)
		return
	}

	// 結果を表示する
	for _, item := range rss.Channel.Items {
		fmt.Println("{")
		fmt.Println("  title:", "\""+item.Title+"\",")
		fmt.Println("  link:", "\""+item.Link+"\",")
		fmt.Println("  guid:", "\""+item.GUID+"\",")
		fmt.Println("  pubDate:", "\""+item.PubDate+"\",")
		fmt.Println("  description:", "\""+item.Description+"\",")
		fmt.Println("},")
	}
}
