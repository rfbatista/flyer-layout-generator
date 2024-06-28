import axios from "axios";

const apiClient = axios.create({
  baseURL: '/api/',
  timeout: 1000,
  headers: {'X-Custom-Header': 'foobar'}
});

export { apiClient }
