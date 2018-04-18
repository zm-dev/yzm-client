import axios from 'axios';

const http = axios.create({
  baseURL: '/api',
  timeout: 6000,
  responseType: 'json',
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  },
});

export default http;
