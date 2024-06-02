import AsyncStorage from "@react-native-async-storage/async-storage";
import { chatAPI } from "./config";
import { Socket, io } from "socket.io-client";

interface SocketIO {
  socket: Socket | null;
}

export const socketIO: SocketIO = {
  socket: null,
};

export const setSocket = async (access_token: string) => {
  try {
    await AsyncStorage.setItem("access_token", access_token);
    socketIO.socket = io(chatAPI, {
      transports: ["websocket"],
      auth: {
        access_token: access_token,
      },
    });
  } catch (e) {
    console.error(e);
  }
};
