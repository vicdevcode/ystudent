import { Button, Text } from "@rneui/base";
import { View } from "react-native";
import TestChat from "../../components/chat/Test";
import { api } from "../../lib/config";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { useEffect, useState } from "react";
import { setSocket, socketIO } from "../../lib/socket";

const ChatScreen = () => {
  const [chats, setChats] = useState([]);
  const [users, setUsers] = useState<string[]>([]);

  const getChats = async () => {
    const token = await AsyncStorage.getItem("access_token");
    try {
      const res = await fetch(api + "/chat/get-all", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      const json = await res.json();
      if (res.status != 200) return;
      setChats(json["chats"]);
    } catch (e) {
      console.error(e);
    }
  };

  const joinRoom = (name: string) => {
    socketIO.socket?.emit("joinRoom", name);
  };

  useEffect(() => {
    // getChats();
    if (socketIO.socket) {
      socketIO.socket.io.on("open", () => console.log("wtf"));
      socketIO.socket.io.on("close", () => console.log("no"));
      socketIO.socket.on("userJoined", (data) => {
        setUsers([...users, data.userId as string]);
      });
    }

    return () => {
      if (socketIO.socket) {
        socketIO.socket.disconnect();
        socketIO.socket.removeAllListeners();
      }
    };
  }, []);

  return (
    <View>
      <Text>It is chat</Text>
      {chats.map((chat) => (
        <Button key={chat["id"]} onPress={() => joinRoom(chat["name"])}>
          <Text>{chat["name"]}</Text>
        </Button>
      ))}
      {users.map((user) => (
        <Text key={user}>{user}</Text>
      ))}
    </View>
  );
};

export default ChatScreen;
