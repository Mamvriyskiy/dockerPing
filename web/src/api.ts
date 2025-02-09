import axios from 'axios';

const API_URL = 'http://localhost:8000/api/ping';

export const sendPingData = async (data: any) => {
  try {
    const response = await axios.post(API_URL, data);
    return response.data;
  } catch (error) {
    console.error('Error sending ping data:', error);
    throw error;
  }
};
