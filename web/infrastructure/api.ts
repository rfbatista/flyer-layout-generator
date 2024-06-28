import axios from "axios";

const apiClient = axios.create({
  baseURL: '/api/',
  headers: {'X-Custom-Header': 'foobar'}
});

export { apiClient }
