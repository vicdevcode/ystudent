import { chatAPI } from "./config";
import { Socket, io } from "socket.io-client";

interface SocketIO {
  socket: Socket | null;
}

export const socketIO: SocketIO = {
  socket: null,
};

export const setSocket = (access_token: string) => {
  socketIO.socket = io(chatAPI, {
    transports: ["websocket"],
    auth: {
      access_token: access_token,
    },
  });
};
