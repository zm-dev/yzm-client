import axios from 'axios';

export default http = axios.create({
  baseURL: '/api',
  timeout: 6000,
  responseType: 'json',
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  },
});
