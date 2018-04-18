import axios from 'axios';

const http = axios.create({
  baseURL: '/api',
  responseType: 'json',
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  },
});

export default http;
