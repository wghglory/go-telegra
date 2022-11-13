import axios from 'axios';

const http = axios.create({
  // https://github.com/vitejs/vite/issues/1149 Jest testing issues.
  baseURL: `${import.meta.env.VITE_APP_HOST}/${import.meta.env.VITE_APP_PREFIX}`,
  // withCredentials: true,
  // timeout: 10000,
});

export default http;
