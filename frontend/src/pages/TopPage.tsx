// TopPage.tsx
import React from 'react';
import { useNavigate } from 'react-router-dom';

const TopPage: React.FC = () => {
  const navigate = useNavigate();

  const goToSelection = () => {
    navigate('/selection');
  };

  return (
    <div>
      <h1>Welcome to g-now!!</h1>
      <button onClick={goToSelection}>詳しくジャンルを見る</button>
    </div>
  );
}

export default TopPage;