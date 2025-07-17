import axios from "axios";
import { v4 as uuidv4 } from "uuid";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

api.interceptors.request.use((config) => {
  const requestId = uuidv4();
  const timestamp = new Date().toISOString();

  config.headers["Request-ID"] = `${requestId}-${timestamp}`;

  if (!config.url.includes("/login") && !config.url.includes("/register")) {
    const token = localStorage.getItem("authToken");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
  }

  return config;
});


export default api;
