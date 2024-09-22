// src/pages/SelectionPage.tsx
import React from 'react';
import { useNavigate } from 'react-router-dom';

const SelectionPage: React.FC = () => {
  const navigate = useNavigate();

  const handleSelection = (category: string) => {
    // Navigate to the genre page with the selected genre as a route parameter
    navigate(`/category/${category}`);
  };

  return (
    <div>
      <h1>Select a Category</h1>
      <button onClick={() => handleSelection('action')}>Action</button>
      <button onClick={() => handleSelection('comedy')}>Comedy</button>
      <button onClick={() => handleSelection('drama')}>Drama</button>
      {/* Add more genres as needed */}
    </div>
  );
}

export default SelectionPage;