import React from 'react';
import { useNavigate } from 'react-router-dom';
import './style.css'; // CSSファイルをインポート

const SelectionPage: React.FC = () => {
  const navigate = useNavigate();

  const handleSelection = (category: string) => {
    navigate(`/category/${category}`);
  };

  return (
    <div className="select-container">
      <h1 className="select-title">Select a Category</h1>
      <button className='select-button' onClick={() => handleSelection('1')}>政治</button>
      <button className='select-button' onClick={() => handleSelection('2')}>エンタメ</button>
      <button className='select-button' onClick={() => handleSelection('3')}>経済</button>
      <button className='select-button' onClick={() => handleSelection('4')}>スポーツ</button>
      <button className='select-button' onClick={() => handleSelection('5')}>ビジネス</button>
      <button className='select-button' onClick={() => handleSelection('6')}>天気</button>
      <button className='select-button' onClick={() => handleSelection('7')}>キャリア</button>
      <button className='select-button' onClick={() => handleSelection('8')}>テクノロジー</button>
      <button className='select-button' onClick={() => handleSelection('9')}>ゲーム</button>
      <button className='select-button' onClick={() => handleSelection('10')}>グルメ</button>
      <button className='select-button' onClick={() => handleSelection('11')}>Web3</button>
    </div>
  );
}

export default SelectionPage;