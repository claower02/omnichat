import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || '';

export const client = axios.create({
  baseURL: `${API_URL}/api/v1`,
});

export const getWSUrl = () => {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const host = window.location.host;
  return `${protocol}//${host}/ws`;
};
