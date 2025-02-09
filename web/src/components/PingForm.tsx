import React, { useState } from 'react';
import axios from 'axios';

const PingForm: React.FC = () => {
  const [containerId, setContainerId] = useState('');
  const [status, setStatus] = useState('');
  const [ipAddress, setIpAddress] = useState(''); // Для ввода IP-адреса

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const token = localStorage.getItem('jwt');
    if (!token) {
        throw new Error('Token not found');
    }

    try {
        const data = { containerId, status, ipAddress };  // Добавили ipAddress
        console.log(token)
        const response = await axios.post('http://localhost:8000/api/ping', data, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        });
      console.log('Ping data sent successfully:', response.data);
      setContainerId('');
      setStatus('');
      setIpAddress('');  // Очищаем поля формы после отправки
      alert('Ping data sent successfully');
    } catch (error) {
      console.error('Error sending ping data:', error);
      alert('Error sending ping data');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>IP Address:</label>
        <input
          type="text"
          value={ipAddress}
          onChange={(e) => setIpAddress(e.target.value)}  // Обработчик для поля IP
          required
        />
      </div>
      <button type="submit">Send Ping</button>
    </form>
  );
};

export default PingForm;
