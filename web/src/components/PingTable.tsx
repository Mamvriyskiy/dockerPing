import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface PingData {
  containerId: string;
  timestamp: string;
  status: string;
}

const PingTable: React.FC = () => {
  const [pingData, setPingData] = useState<PingData[]>([]);
  
  const fetchPingData = async () => {
    try {
      const response = await axios.get('http://localhost:8000/api/ping');
      setPingData(response.data);
    } catch (error) {
      console.error('Error fetching ping data:', error);
    }
  };

  useEffect(() => {
    fetchPingData();
    const interval = setInterval(fetchPingData, 5000); // Обновляем данные каждые 5 секунд
    return () => clearInterval(interval); // Очистка интервала при размонтировании компонента
  }, []);

  return (
    <div>
      <h2>Ping Data Table</h2>
      <table>
        <thead>
          <tr>
            <th>Container ID</th>
            <th>Timestamp</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {pingData.map((data) => (
            <tr key={data.containerId}>
              <td>{data.containerId}</td>
              <td>{data.timestamp}</td>
              <td>{data.status}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PingTable;
