import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'; // Используйте Routes вместо Switch
import PingForm from './PingForm.tsx';
import LoginPage from './LoginPage.tsx';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/ping" element={<PingForm />} />
        <Route path="/" element={<h1>Welcome! Please <a href="/login">Login</a></h1>} />
      </Routes>
    </Router>
  );
};

export default App;
