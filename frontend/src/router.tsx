// src/router.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { TopPage, SelectionPage, CategoryPage } from './pages/pages';

const AppRouter: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<TopPage />} />
        <Route path="/selection" element={<SelectionPage />} />
        {/* Dynamic route to handle different genres */}
        <Route path="/category/:id" element={<CategoryPage />} />
      </Routes>
    </Router>
  );
}

export default AppRouter;