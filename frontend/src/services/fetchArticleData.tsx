const apiUrl = "http://localhost:8080";

interface Article {
  categories: Array<{ category_name: string }>;
  tags: Array<{ tag_name: string }>;
}

interface GraphData {
  nodes: Array<{ id: string }>;
  links: Array<{ source: string; target: string }>;
}

// APIからデータを取得し、グラフ用データに変換する関数
export async function fetchArticlesData(categoryNumber: string): Promise<GraphData> {
  try {
    const endpoint = `${apiUrl}/categories/${categoryNumber}/articles`
    const response = await fetch(endpoint);
    if (!response.ok) {
      throw new Error('Failed to fetch articles data');
    }

    const articles: Article[] = await response.json();

    // ノードとリンクを初期化
    const nodes: Set<string> = new Set();  // 重複しないようにSetを使用
    const links: Array<{ source: string; target: string }> = [];

    for(let i: number = 0; i < articles.length; i++) {
      console.log(articles[0])
      articles[i].categories.forEach(category => {
            nodes.add(category.category_name);
            articles[i].tags.forEach(tag => {
              nodes.add(tag.tag_name);
              links.push({ source: category.category_name, target: tag.tag_name });
            });
      });
    }

    // ノードを配列に変換（Setから配列に変換）
    const nodeArray = Array.from(nodes).map(node => ({ id: node }));

    // グラフ用のデータを返す
    return {
      nodes: nodeArray,
      links,
    };
  } catch (error) {
    console.error('Error fetching articles data:', error);
    throw error;
  }
}