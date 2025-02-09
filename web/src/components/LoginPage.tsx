import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const LoginPage: React.FC = () => {
  const navigate = useNavigate(); // Получаем navigate для перенаправления
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null); // Для отображения ошибок

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    // Данные для отправки на сервер
    const credentials = {
      email,
      password
    };

    try {
      // Отправляем запрос на сервер
      const response = await fetch('http://localhost:8000/auth/sign-in', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
      });

      if (!response.ok) {
        throw new Error('Неверные данные для авторизации');
      }

      // Получаем токен из ответа
      const data = await response.json();
      const token = data.token; // Здесь предполагается, что сервер возвращает объект с токеном

      // Сохраняем токен в localStorage (или в состояние, если требуется)
      localStorage.setItem('authToken', token);

      // Переходим на страницу после успешной авторизации
      navigate('/ping');
    } catch (err) {
      setError(err.message); // Устанавливаем ошибку, если она возникла
    }
  };

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleLogin}>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
        />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
        />
        <button type="submit">Login</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>} {/* Отображаем ошибку, если она есть */}
    </div>
  );
};

export default LoginPage;
