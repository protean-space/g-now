import React from 'react';
import { useNavigate } from 'react-router-dom';
import './style.css'; // CSSファイルをインポート

const TopPage: React.FC = () => {
  const navigate = useNavigate();

  const goToSelection = () => {
    navigate('/selection');
  };

  return (
    <div className="toppage-container">
      <h1 className="toppage-title">Welcome to g-now!!</h1>
      <button onClick={goToSelection} className="goToSelection">詳しくジャンルを見る</button>
    </div>
  );
}

export default TopPage;