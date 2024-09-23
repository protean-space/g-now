import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Graph from '../components/Graph';
import Modal from '../components/Modal'; // モーダルコンポーネントをインポート
import { fetchArticlesData } from '../services/fetchArticleData';
import { fetchCategoryData } from '../services/fetchCategoryData';
import { fetchRelatedNews } from '../services/fetchRelatedNews';

const CategoryPage: React.FC = () => {
  const { id } = useParams<{ id: string }>(); // URLからカテゴリIDを取得
  const [graphData, setGraphData] = useState(null);
  const [categoryData, setCategoryData] = useState(null);
  const [relatedNews, setRelatedNews] = useState([]); // クリックしたノードに関連するニュースを格納
  const [loading, setLoading] = useState(true);
  const [isModalOpen, setModalOpen] = useState(false); // モーダルの開閉状態を管理

  useEffect(() => {
    if (id) {
      // 非同期でなんとかする
      const fetchData = async () => {
        try {
          setLoading(true);
          const [articles, category] = await Promise.all([
            fetchArticlesData(id), // 記事データの取得
            fetchCategoryData(id)   // カテゴリデータの取得
          ]);

          setGraphData(articles);
          setCategoryData(category);
        } catch (error) {
          console.error("Error fetching data:", error);
        } finally {
          setLoading(false);
        }
      };

      fetchData();
    }
  }, [id]);

  // 関連ニュース取得関数
  const handleFetchRelatedNews = async (nodeId: string) => {
    try {
      const news = await fetchRelatedNews(nodeId); // クリックされたノードIDに関連するニュースを取得
      setRelatedNews(news); // ニュースデータを状態にセット
      setModalOpen(true);   // モーダルを表示
    } catch (error) {
      console.error("Error fetching related news:", error);
    }
  };

  const closeModal = () => {
    setModalOpen(false); // モーダルを閉じる
  };

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <div>
      {categoryData ? <h3>{categoryData}</h3> : <p>No Category Data</p>}
      {graphData ? <Graph data={graphData} onNodeClick={handleFetchRelatedNews} /> : <p>No Articles Data</p>}

      {/* モーダルを表示 */}
      <Modal isOpen={isModalOpen} onClose={closeModal} relatedNews={relatedNews} />
    </div>
  );
};

export default CategoryPage;