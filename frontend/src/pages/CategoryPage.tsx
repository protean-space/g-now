import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import Graph from '../components/Graph';
import { fetchCategoryData } from '../services/fetchCategoryData';

const CategoryPage: React.FC = () => {
  const { categoryName } = useParams<{ categoryName: string }>();
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    // Fetch data based on the category
    if (categoryName) {
      fetchCategoryData(categoryName).then((fetchedData) => {
        setData(fetchedData);
      });
    }
  }, []);

  return (
    <div>
      
      <h1>{data ? data: "Not Category Data Found"}</h1>
      {data ? (
        <div>
          {/* Use the common Graph component to render the data */}
          <Graph data={data} />
        </div>
      ) : (
        <p>Loading data...</p>
      )}
    </div>
  );
};

export default CategoryPage;