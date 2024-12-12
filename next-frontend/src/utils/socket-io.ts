import { io } from "socket.io-client";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;
console.log('apiUrl', apiUrl);
export const socket = io(`${apiUrl}`, {
  autoConnect: false,
});