import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'; // Используйте Routes вместо Switch
import PingTable from './components/PingTable.tsx';
import LoginPage from './components/LoginPage.tsx';
import RegisterPage from './components/RegisterPage.tsx';
import styles from './App.module.css';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/sign-in" element={<LoginPage />} />
        <Route path="/sign-up" element={<RegisterPage />} />
        <Route path="/ping" element={<PingTable />} />
        <Route path="/" element={
          <div className={styles.centeredContainer}>
            <h1>
              Добро пожаловать!
            </h1>
            <h1> 
              <a href="/sign-in">Войдите</a> или <a href="/sign-up">Зарегистрируйтесь</a>
            </h1>
          </div>
        } />
      </Routes>
    </Router>
  );
};

export default App;

// <div className="centered-container">
//       <h1>
//         Welcome! Please <Link to="/login">Login</Link> or <Link to="/register">Register</Link>
//       </h1>
//     </div>
