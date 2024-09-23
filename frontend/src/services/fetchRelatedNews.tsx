/**
 * fetching articles by a tag
 */

export async function fetchRelatedNews(tagName: string) {
    const apiUrl = "http://localhost:8080";
    const response = await fetch(`${apiUrl}/tags/${tagName}/articles`); // nodeIdに基づいてAPIリクエストを送信
    console.log(response)
    const data = await response.json();
    console.log(data)
    return data; // 関連ニュースデータを返す
  }