import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Graph from '../components/Graph';
import { fetchArticlesData } from '../services/fetchArticleData';
import { fetchCategoryData } from '../services/fetchCategoryData';

const CategoryPage: React.FC = () => {
  const { id } = useParams<{ id: string }>(); // URLからカテゴリIDを取得
  const [graphData, setGraphData] = useState(null);
  const [categoryData, setCategoryData] = useState(null);
  const [loading, setLoading] = useState(true);

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

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <div>
      {categoryData ? <h2>{categoryData}</h2> : <p>No Category Data</p>}
      {graphData ? <Graph data={graphData} /> : <p>No Articles Data</p>}
    </div>
  );
};

export default CategoryPage;