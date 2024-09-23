import React from 'react';
import AppRouter from './router';
import Header from './components/Header'

const App: React.FC = () => {
  return (
    <div>
      <Header />
      <AppRouter />
    </div>
  );
}

export default App;