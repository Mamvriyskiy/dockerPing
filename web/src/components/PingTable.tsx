import React, { useEffect, useState } from 'react';
import axios from 'axios';
import styles from './PingTable.module.css';

interface PingData {
  ipcontainer: string;
  timeping: string;
  status: string;
}

const PingTable: React.FC = () => {
  const [pingData, setPingData] = useState<PingData[]>([]);

  const fetchPingData = async () => {
    const token = localStorage.getItem('jwt');
    console.log("==", token)
    if (!token) {
      throw new Error('Token not found');
    }

    console.log("==", token)
    try {
      const response = await axios.get('http://localhost:8000/api/pinger', {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token
        }
      });
      console.log(response.data)
      setPingData(response.data);
    } catch (error) {
  
      console.error('Error fetching ping data:', error);
    }
  };

  useEffect(() => {
    fetchPingData();
    const interval = setInterval(fetchPingData, 5000);
    return () => clearInterval(interval);
  }, []);

  const formatTimestamp = (timestamp: string) => {
    const date = new Date(timestamp);
    
    return date.toLocaleString('en-GB', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      fractionalSecondDigits: 3,
    });
  };

  return (
    <div className={styles.container}>
      <div className={styles.tableBox}>
        <h2 className={styles.title}>Таблица состояния контейнеров</h2>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>IP контейнера</th>
              <th>Время и дата</th>
              <th>Состояние</th>
            </tr>
          </thead>
          <tbody>
            {pingData && pingData.length > 0 ? (
              pingData.map((data) => (
                <tr key={data.containerid}>
                  <td>{data.ipcontainer}</td>
                  <td>{formatTimestamp(data.timeping)}</td>
                  <td>{data.status}</td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan={3}>No ping data available</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
      <PingForm />
    </div>
  );
};

const PingForm: React.FC = () => {
  const [ipcontainer, setIpAddress] = useState('');
  const [error, setError] = useState('');

  const isValidIP = (ip: string): boolean => {
    const regex = /^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$/;
    return regex.test(ip);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    if (!isValidIP(ipcontainer)) {
      setError('IP-адрес введен неверно');
      return;
    }

    setError('');

    const token = localStorage.getItem('jwt');
    if (!token) {
      throw new Error('Token not found');
    }

    try {
      const data = { ipcontainer };
      console.log(token);
      const response = await axios.post('http://localhost:8000/api/ping', data, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token
        }
      });
      console.log('Ping data sent successfully:', response.data);
      setIpAddress('');
      alert('Ping data sent successfully');
    } catch (error) {
      console.error('Error sending ping data:', error);
      alert('Error sending ping data');
    }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <div>
        <label>IP-адрес контейнера:</label>
        <input
          type="text"
          value={ipcontainer}
          onChange={(e) => setIpAddress(e.target.value)}
          required
        />
        {error && <div style={{ color: 'red' }}>{error}</div>}
      </div>
      <button type="submit">Добавить контейнер</button>
    </form>
  );
};

export default PingTable;
