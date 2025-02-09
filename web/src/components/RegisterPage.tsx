import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './SignPage.module.css'; // Подключаем стили

const SignUpPage: React.FC = () => {
  const navigate = useNavigate();
  const [login, setLogin] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleSignUp = async (e: React.FormEvent) => {
    e.preventDefault();

    const userData = { login, email, password };

    try {
      const response = await fetch('http://backend:8000/auth/sign-up', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userData),
      });

      if (!response.ok) throw new Error('Ошибка регистрации');

      navigate('/sign-in');
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.signupBox}>
        <h1 className={styles.title}>Регистрация</h1>
        <form onSubmit={handleSignUp} className={styles.form}>
          <input
            type="text"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            placeholder="Введите логин"
            className={styles.input}
          />
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
          <button type="submit" className={styles.button}>Зарегистрироваться</button>
        </form>
        {error && <p className={styles.error}>{error}</p>}
        <p className={styles.loginLink}>
          Уже есть аккаунт? <a href="/sign-in">Войдите</a>
        </p>
      </div>
    </div>
  );
};

export default SignUpPage;
