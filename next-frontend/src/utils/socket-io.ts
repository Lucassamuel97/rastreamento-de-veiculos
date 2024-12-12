import { io } from "socket.io-client";

// Tem que ser Localhost porque o navegador não conecta diretamente com o api:300 
export const socket = io("http://localhost:3000", {
  autoConnect: false,
});