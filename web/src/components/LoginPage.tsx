import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './SignPage.module.css'; // Импортируем CSS-модуль

const LoginPage: React.FC = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    const credentials = { email, password };

    try {
      const response = await fetch('http://localhost:8000/auth/sign-in', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials),
      });

      if (!response.ok) throw new Error('Неверные данные для авторизации');

      const data = await response.json();
      localStorage.setItem('jwt', data.Token);
      navigate('/ping');
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.loginBox}>
        <h1 className={styles.title}>Вход</h1>
        <form onSubmit={handleLogin} className={styles.form}>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Введите email"
            className={styles.input}
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Введите пароль"
            className={styles.input}
          />
          <button type="submit" className={styles.button}>Войти</button>
        </form>
        {error && <p className={styles.error}>{error}</p>}
        <p className={styles.registerLink}>
          Нет аккаунта? <a href="/sign-up">Зарегистрируйтесь</a>
        </p>
      </div>
    </div>
  );
};

export default LoginPage;
