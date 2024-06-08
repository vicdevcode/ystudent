import { Button, CheckBox, Dialog, Icon, Input, Text } from "@rneui/base";
import { Pressable, StyleSheet, View } from "react-native";
import { FC, useEffect, useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import {
  isoStringToDate,
  parseJwt,
  roleTypes,
  shortFio,
} from "../../lib/utils";
import { ScrollView } from "react-native-gesture-handler";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { chatAPI } from "../../lib/config";
import { Socket, io } from "socket.io-client";

interface User {
  id: string;
  email: string;
  firstname: string;
  surname: string;
  middlename: string;
  roleType: string;
}

interface ChatScreenProps {
  id: string;
  name: string;
  members: User[];
  messages: Message[];
  token: string;
  type: string;
  back: () => void;
}

interface Message {
  id?: string;
  senderId?: string;
  content: string;
  senderFio: string;
  createdAt: string;
  important: boolean;
}

const ChatScreen: FC<ChatScreenProps> = ({
  id,
  name,
  messages,
  members,
  back,
  token,
  type,
}) => {
  const [chatMessages, setMessages] = useState<Message[]>(messages);
  const [message, setMessage] = useState<string>("");
  const [candidates, setCandidates] = useState<User[]>([]);
  const [filteredMessages, setFilteredMessages] = useState<Message[]>(messages);
  const [showMembers, setShowMembers] = useState(false);
  const [showCandidates, setShowCandidates] = useState(false);
  const [showChangeImportant, setShowChangeImportant] = useState(false);
  const [checkImportant, setCheckImportant] = useState(false);
  const [messageId, setMessageId] = useState<string>("");
  const [socket, setSocket] = useState<Socket | null>(null);
  const [isUserEditor, setIsUserEditor] = useState<boolean>(false);
  const [currentMembers, setCurrentMembers] = useState<User[]>(members);

  const toggleShowMembers = () => {
    setShowMembers(!showMembers);
  };

  const toggleShowCandidates = () => {
    setShowCandidates(!showCandidates);
  };

  const toggleShowChangeImportant = (id: string, important: boolean) => {
    setMessageId(id);
    setCheckImportant(important);
    setShowChangeImportant(!showChangeImportant);
  };

  const isEditor = async () => {
    if (!(type == "NEWS" || type == "USER_NEWS")) return setIsUserEditor(true);
    try {
      const token = await AsyncStorage.getItem("access_token");
      const payload = parseJwt(token as string);
      const res = await fetch(chatAPI + "/api/v1/chat/check-admin/", {
        method: "POST",
        headers: {
          Authorization: "Bearer " + token,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          chat_id: id,
          user_id: payload["id"],
        }),
      });
      const json = await res.json();
      if (res.status != 200) return;
      if (!json) return;
      setIsUserEditor(true);
    } catch (e) {
      console.error(e);
      setIsUserEditor(false);
    }
  };

  const getChatsMessages = async () => {
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
        if (json["chats"][i]["id"] != id) continue;
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
        setMessages(json["chats"][i]["messages"]);
        setFilteredMessages(json["chats"][i]["messages"]);
      }
    } catch (e) {
      console.error(e);
    }
  };

  const changeImportant = async () => {
    const token = await AsyncStorage.getItem("access_token");
    const res = await fetch(chatAPI + "/api/v1/chat/important/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: messageId,
        important: !checkImportant,
        is_news: false,
        is_task: false,
      }),
    });
    if (res.status !== 200) return;
    setCheckImportant(!checkImportant);
    await getChatsMessages();
  };

  const sendMessage = async () => {
    if (!socket) return;
    socket.emit("send_message", message);
    setTimeout(async () => {
      await getChatsMessages();
      setMessage("");
    }, 500);
  };

  const getCandidates = async () => {
    const token = await AsyncStorage.getItem("access_token");
    const res = await fetch(chatAPI + "/api/v1/chat/candidates/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: id,
      }),
    });
    const json = await res.json();
    if (res.status !== 200) return;
    setCandidates(json);
  };

  const addMember = async (user_id: string) => {
    const token = await AsyncStorage.getItem("access_token");
    const res = await fetch(chatAPI + "/api/v1/chat/add/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: id,
        user_id: user_id,
      }),
    });
    const json = await res.json();
    console.log(json);
    if (res.status !== 200) return;
    setCurrentMembers([...currentMembers, json]);
    await getCandidates();
  };

  const sendNews = async () => {
    if (!socket) return;
    socket.emit("send_news", message);
    setTimeout(async () => {
      await getChatsMessages();
      setMessage("");
    }, 500);
  };

  const filterMessages = () => {
    if (filteredMessages.length != chatMessages.length) {
      setFilteredMessages(chatMessages);
    } else {
      setFilteredMessages(chatMessages.filter((msg) => msg.important));
    }
  };

  const joinRoom = (good_socket: Socket, name: string) => {
    if (!good_socket) return;
    good_socket.emit("join_room", name);
  };

  useEffect(() => {
    isEditor();
    getCandidates();
    const good_socket = io(chatAPI, {
      transports: ["websocket"],
      auth: {
        access_token: token,
      },
    });

    setSocket(good_socket);

    getChatsMessages();
    joinRoom(good_socket, id);
    if (good_socket) {
      good_socket.on("receive_message", () => {
        getChatsMessages();
      });
    }
    return () => {
      if (!good_socket) return;
      good_socket.emit("leave_room");
      good_socket.removeAllListeners();
      good_socket.disconnect();
    };
  }, []);

  return (
    <SafeAreaView
      style={{
        flex: 1,
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <View style={styles.header}>
        <Icon name="chevron-left" type="material" size={32} onPress={back} />
        <Text
          style={{
            fontSize: 16,
          }}
        >
          {name}
        </Text>
        {!(type == "NEWS" || type == "USER_NEWS") && (
          <>
            <Icon
              name="people"
              type="material"
              size={24}
              onPress={toggleShowMembers}
            />
            <Icon
              name="feedback"
              type="material"
              size={24}
              onPress={filterMessages}
            />
            {type == "CUSTOM" && (
              <Icon
                name="add"
                type="material"
                size={24}
                onPress={toggleShowCandidates}
              />
            )}
          </>
        )}
      </View>
      {!(type == "NEWS" || type == "USER_NEWS") && (
        <>
          <Dialog
            isVisible={showMembers}
            onBackdropPress={toggleShowMembers}
            overlayStyle={{
              backgroundColor: "#fff",
            }}
            statusBarTranslucent
          >
            <Dialog.Title title="Участники" />
            <ScrollView
              style={{
                maxHeight: 400,
              }}
            >
              {currentMembers.map((member) => (
                <View
                  style={{
                    marginBottom: 4,
                    paddingHorizontal: 12,
                    paddingVertical: 4,
                    borderRadius: 12,
                    backgroundColor: "#eee",
                  }}
                  key={member.id}
                >
                  <Text style={{ fontSize: 14 }}>
                    {shortFio(
                      `${member.surname} ${member.firstname} ${member.middlename}`,
                    )}
                  </Text>
                  {(member.roleType === "TEACHER" ||
                    member.roleType === "STUDENT") && (
                      <Text style={{ color: "#999", fontSize: 12 }}>
                        {roleTypes[member.roleType]}
                      </Text>
                    )}
                </View>
              ))}
            </ScrollView>
          </Dialog>
          <Dialog
            isVisible={showChangeImportant}
            onBackdropPress={() => setShowChangeImportant(!showChangeImportant)}
            overlayStyle={{
              backgroundColor: "#fff",
            }}
            statusBarTranslucent
          >
            <Dialog.Title title="Важность" />
            <CheckBox
              title="Это сообщение важное?"
              checked={checkImportant}
              onPress={changeImportant}
            />
          </Dialog>
          {type == "CUSTOM" && (
            <Dialog
              isVisible={showCandidates}
              onBackdropPress={toggleShowCandidates}
              overlayStyle={{
                backgroundColor: "#fff",
              }}
              statusBarTranslucent
            >
              <Dialog.Title title="Добавить участника" />
              <ScrollView
                style={{
                  maxHeight: 400,
                }}
              >
                {candidates.map((candidate) => (
                  <Pressable
                    style={{
                      marginBottom: 4,
                      paddingHorizontal: 12,
                      paddingVertical: 4,
                      borderRadius: 12,
                      backgroundColor: "#eee",
                    }}
                    key={candidate.id}
                    onPress={() => addMember(candidate.id)}
                  >
                    <Text style={{ fontSize: 14 }}>
                      {shortFio(
                        `${candidate.surname} ${candidate.firstname} ${candidate.middlename}`,
                      )}
                    </Text>
                    {(candidate.roleType === "TEACHER" ||
                      candidate.roleType === "STUDENT" ||
                      candidate.roleType === "EMPLOYEE") && (
                        <Text style={{ color: "#999", fontSize: 12 }}>
                          {roleTypes[candidate.roleType]}
                        </Text>
                      )}
                  </Pressable>
                ))}
              </ScrollView>
            </Dialog>
          )}
        </>
      )}
      <ScrollView
        contentContainerStyle={{
          justifyContent: "flex-end",
          flexDirection: "column",
          flex: 1,
        }}
      >
        <ScrollView
          contentContainerStyle={{
            justifyContent: "flex-end",
            alignItems: "flex-start",
            paddingHorizontal: 8,
            gap: 8,
          }}
        >
          {filteredMessages &&
            filteredMessages.map((msg, i) => (
              <Pressable
                onLongPress={() =>
                  isUserEditor &&
                  toggleShowChangeImportant(msg.id as string, msg.important)
                }
                key={i}
                style={{ width: "100%" }}
              >
                <View
                  style={
                    msg.senderFio === "me"
                      ? {
                        borderRadius: 12,
                        backgroundColor: "#e9e9e9",
                        justifyContent: "flex-end",
                        alignItems: "flex-end",
                        width: "100%",
                        paddingHorizontal: 16,
                        paddingVertical: 12,
                      }
                      : {
                        borderRadius: 12,
                        backgroundColor: "#ddd",
                        paddingHorizontal: 16,
                        paddingVertical: 12,
                      }
                  }
                >
                  {msg.senderFio !== "me" && (
                    <Text style={{ fontSize: 12 }}>
                      {shortFio(msg.senderFio)}
                    </Text>
                  )}
                  <Text style={{ fontSize: 16 }}>{msg.content}</Text>
                  <Text
                    style={{
                      fontSize: 10,
                      width: "100%",
                      textAlign: msg.senderFio === "me" ? "left" : "right",
                    }}
                  >
                    {isoStringToDate(msg.createdAt)}
                  </Text>
                </View>
              </Pressable>
            ))}
        </ScrollView>
        {isUserEditor && (
          <View
            style={{
              flexDirection: "row",
              marginTop: 12,
            }}
          >
            <View
              style={{
                flex: 1,
              }}
            >
              <Input
                value={message}
                onChangeText={(value) => setMessage(value)}
              />
            </View>
            <Button
              onPress={type == "NEWS" ? sendNews : sendMessage}
              icon={{
                name: "send",
                type: "material",
                color: "#FFF",
              }}
              buttonStyle={{
                borderRadius: 12,
              }}
              containerStyle={{
                width: 60,
                paddingRight: 8,
              }}
            />
          </View>
        )}
      </ScrollView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  header: {
    alignItems: "center",
    justifyContent: "flex-start",
    flexDirection: "row",
    height: 50,
    padding: 8,
    gap: 12,
    borderBottomWidth: 1,
    marginBottom: 12,
  },
});

export default ChatScreen;
