import { chatAPI } from "./config";
import { io } from "socket.io-client";

const socket = io(chatAPI, {
  transports: ["websocket"],
  auth: {
    access_token: "",
  },
});

export default socket;
