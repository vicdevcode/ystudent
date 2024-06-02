import AsyncStorage from "@react-native-async-storage/async-storage";
import { useEffect, useState } from "react";
import { View } from "react-native";
import { api } from "../../lib/config";
import { Button, Text } from "@rneui/base";
import ChatScreen from "./ChatScreen";
import { SafeAreaView } from "react-native-safe-area-context";
import { parseJwt } from "../../lib/utils";

interface Chat {
  id: string;
  name: string;
  members: {
    id: string;
    email: string;
    firstname: string;
    surname: string;
    middlename: string;
    roleType: string;
  }[];
  messages: Message[];
  type: string;
}

interface Message {
  senderId?: string;
  content: string;
  senderFio: string;
  createdAt: string;
}

const ChatListScreen = () => {
  const [chats, setChats] = useState<Chat[]>([]);
  const [selectedChat, setSelectedChat] = useState<Chat | null>(null);

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
      const payload = parseJwt(token as string);

      for (let i = 0; i < json["chats"].length; i++) {
        for (let j = 0; j < json["chats"][i]["messages"].length; j++) {
          if (json["chats"][i]["messages"][j]["senderId"] === payload["id"]) {
            json["chats"][i]["messages"][j]["senderFio"] = "me";
            console.log("me");
          }
        }
      }
      setChats(json["chats"]);
      console.log(json["chats"]);
    } catch (e) {
      console.error(e);
    }
  };

  const back = () => {
    setSelectedChat(null);
  };

  useEffect(() => {
    getChats();
  }, []);

  if (selectedChat != null)
    return (
      <ChatScreen
        id={selectedChat.id}
        name={selectedChat.name}
        messages={selectedChat.messages}
        members={selectedChat.members}
        back={back}
        chats={getChats}
      />
    );

  return (
    <SafeAreaView>
      <View
        style={{
          height: 50,
          padding: 8,
          alignItems: "center",
          justifyContent: "flex-start",
          flexDirection: "row",
          borderBottomWidth: 1,
          marginBottom: 16,
        }}
      >
        <Text style={{ fontSize: 16 }}>Ваши чаты</Text>
      </View>
      <View style={{ paddingHorizontal: 8, gap: 12 }}>
        {chats.length > 0 ? (
          chats.map((chat) => (
            <Button
              buttonStyle={{
                backgroundColor: "transparent",
                alignItems: "flex-start",
                justifyContent: "flex-start",
                flexDirection: "column",
                paddingHorizontal: 12,
                paddingVertical: 8,
                margin: 0,
                borderWidth: 1,
                borderRadius: 8,
              }}
              onPress={() => setSelectedChat(chat)}
              key={chat.id}
            >
              <Text style={{ fontSize: 20 }}>{chat.name}</Text>
              <Text>Официальная группа</Text>
            </Button>
          ))
        ) : (
          <Text style={{ textAlign: "center" }}>нет чатов</Text>
        )}
      </View>
    </SafeAreaView>
  );
};

export default ChatListScreen;
