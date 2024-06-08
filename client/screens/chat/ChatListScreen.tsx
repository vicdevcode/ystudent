import AsyncStorage from "@react-native-async-storage/async-storage";
import { useEffect, useState } from "react";
import { View } from "react-native";
import { chatAPI } from "../../lib/config";
import { Button, Dialog, Icon, Input, ListItem, Text } from "@rneui/base";
import ChatScreen from "./ChatScreen";
import { SafeAreaView } from "react-native-safe-area-context";
import { parseJwt } from "../../lib/utils";
import { ScrollView } from "react-native-gesture-handler";
import { io } from "socket.io-client";

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

interface AllChatCheckbox {
  id: string;
  name: string;
  type: string;
}

interface Message {
  id?: string;
  senderId?: string;
  content: string;
  senderFio: string;
  createdAt: string;
  important: boolean;
}

const ChatListScreen = () => {
  const [chats, setChats] = useState<Chat[]>([]);
  const [selectedChat, setSelectedChat] = useState<Chat | null>(null);
  const [showSendTo, setShowSendTo] = useState(false);
  const [allChats, setAllChats] = useState<AllChatCheckbox[]>([]);
  const [allChecks, setAllChecks] = useState<boolean[]>([]);
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [role, setRole] = useState<string>("");
  const [token, setToken] = useState<string | null>(null);

  const toogleShowSendTo = () => {
    setShowSendTo(!showSendTo);
  };

  const check = (i: number) => {
    setAllChecks([
      ...allChecks.slice(0, i),
      !allChecks[i],
      ...allChecks.slice(i + 1, allChecks.length),
    ]);
  };

  const getAllChats = async () => {
    const token = await AsyncStorage.getItem("access_token");
    const res = await fetch(chatAPI + "/api/v1/chat/get-all-chats", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    });
    const json = await res.json();
    if (res.status != 200) return;
    const ccc = [];
    for (let i = 0; i < json.length; i++) {
      ccc.push(false);
    }
    setAllChecks(ccc);
    setAllChats(json);
  };

  const getChats = async () => {
    const token = await AsyncStorage.getItem("access_token");
    try {
      const res = await fetch(chatAPI + "/api/v1/chat/get-all", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      });
      const json = await res.json();
      if (res.status != 200) return;
      const payload = parseJwt(token as string);

      for (let i = 0; i < json["chats"].length; i++) {
        json["chats"][i]["messages"] = json["chats"][i]["messages"].sort(
          (a: Message, b: Message) => {
            const adate = Date.parse(a.createdAt);
            const bdate = Date.parse(b.createdAt);
            if (adate < bdate) return -1;
            if (adate > bdate) return 1;
            return 0;
          },
        );
        for (let j = 0; j < json["chats"][i]["messages"].length; j++) {
          if (json["chats"][i]["messages"][j]["senderId"] === payload["id"]) {
            json["chats"][i]["messages"][j]["senderFio"] = "me";
          }
        }
      }
      setChats(json["chats"]);
      setRole(json["roleType"]);
    } catch (e) {
      console.error(e);
    }
  };

  const sendAlertMessage = () => {
    const socket = io(chatAPI, {
      transports: ["websocket"],
      auth: {
        access_token: token,
      },
    });

    for (let i = 0; i < allChecks.length; i++) {
      if (allChecks[i]) {
        console.log(i);
        socket.emit("send_message_to", {
          chat_id: allChats[i]["id"],
          message: alertMessage,
        });
      }
    }
    setTimeout(() => {
      socket.removeAllListeners();
      socket.disconnect();
      setShowSendTo(false);
    }, 1000);
  };

  const setAToken = async () => {
    setToken(await AsyncStorage.getItem("access_token"));
  };

  const back = () => {
    setSelectedChat(null);
  };

  useEffect(() => {
    setAToken();
    getChats();
    getAllChats();
  }, []);

  return (
    <>
      {selectedChat != null ? (
        <ChatScreen
          token={token as string}
          id={selectedChat.id}
          name={selectedChat.name}
          messages={selectedChat.messages}
          members={selectedChat.members}
          back={back}
        />
      ) : (
        <SafeAreaView>
          <View
            style={{
              height: 50,
              padding: 8,
              alignItems: "center",
              justifyContent: "space-between",
              flexDirection: "row",
              borderBottomWidth: 1,
              marginBottom: 16,
            }}
          >
            <Text style={{ fontSize: 16 }}>Ваши чаты</Text>
            {role === "TEACHER" && (
              <Icon
                onPress={toogleShowSendTo}
                name="reply-all"
                type="material"
                style={{
                  marginEnd: 10,
                }}
              />
            )}
          </View>
          <Dialog
            isVisible={showSendTo}
            onBackdropPress={toogleShowSendTo}
            overlayStyle={{
              backgroundColor: "#fff",
              marginBottom: 200,
            }}
            statusBarTranslucent
          >
            <Dialog.Title title="Отправить всем" />
            <ScrollView style={{ maxHeight: 300 }}>
              {allChats.map((achat, i) => (
                <ListItem bottomDivider key={achat.id}>
                  <ListItem.CheckBox
                    iconType="material-community"
                    checkedIcon="checkbox-marked"
                    uncheckedIcon="checkbox-blank-outline"
                    checked={allChecks[i]}
                    onPress={() => check(i)}
                  />
                  <ListItem.Content>
                    <ListItem.Title>{achat.name}</ListItem.Title>
                  </ListItem.Content>
                </ListItem>
              ))}
            </ScrollView>
            <Input
              multiline
              value={alertMessage}
              onChangeText={(value) => setAlertMessage(value)}
            />
            <Button onPress={sendAlertMessage} title="Отправить" />
          </Dialog>
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
                  <Text>{chat.type}</Text>
                </Button>
              ))
            ) : (
              <Text style={{ textAlign: "center" }}>нет чатов</Text>
            )}
          </View>
        </SafeAreaView>
      )}
    </>
  );
};

export default ChatListScreen;
