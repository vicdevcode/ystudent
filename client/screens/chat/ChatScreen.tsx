import { Button, Dialog, Icon, Input, Text } from "@rneui/base";
import { StyleSheet, View } from "react-native";
import { FC, useEffect, useState } from "react";
import { socketIO } from "../../lib/socket";
import { SafeAreaView } from "react-native-safe-area-context";
import { isoStringToDate, roleTypes, shortFio } from "../../lib/utils";
import { ScrollView } from "react-native-gesture-handler";

interface ChatScreenProps {
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
  back: () => void;
  chats: () => Promise<void>;
}

interface Message {
  senderId?: string;
  content: string;
  senderFio: string;
  createdAt: string;
}

const ChatScreen: FC<ChatScreenProps> = ({
  id,
  name,
  messages,
  members,
  back,
  chats,
}) => {
  const [chatMessages, setMessages] = useState<Message[]>(messages);
  const [message, setMessage] = useState<string>("");
  const [showMembers, setShowMembers] = useState(false);

  const toggleShowMembers = () => {
    setShowMembers(!showMembers);
  };

  const sendMessage = () => {
    if (!socketIO.socket) return;
    setMessages([
      ...chatMessages,
      {
        content: message,
        senderFio: "me",
        createdAt: new Date().toISOString(),
      },
    ]);
    socketIO.socket.emit("send_message", message);
    chats();
    setMessage("");
  };
  // const [chats, setChats] = useState([]);
  // const [users, setUsers] = useState<string[]>([]);

  // const getChats = async () => {
  //   const token = await AsyncStorage.getItem("access_token");
  //   try {
  //     const res = await fetch(api + "/chat/get-all", {
  //       method: "GET",
  //       headers: {
  //         Authorization: "Bearer " + token,
  //       },
  //     });
  //     const json = await res.json();
  //     if (res.status != 200) return;
  //     setChats(json["chats"]);
  //   } catch (e) {
  //     console.error(e);
  //   }
  // };

  const joinRoom = (name: string) => {
    if (!socketIO.socket) return;
    socketIO.socket?.emit("join_room", name);
  };

  const leaveRoom = () => {
    if (!socketIO.socket) return;
    socketIO.socket.emit("leave_room");
  };

  useEffect(() => {
    joinRoom(id);
    if (socketIO.socket) {
      socketIO.socket.on("user_joined", (data) => console.log(data));

      socketIO.socket.on("receive_message", (msg: Message) =>
        setMessages([...chatMessages, msg]),
      );
      // socketIO.socket.on("user_joined", (data) => {
      //   setUsers([...users, data.userId as string]);
      // });
    }

    return () => {
      if (socketIO.socket) {
        socketIO.socket.disconnect();
        socketIO.socket.removeAllListeners();
      }
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
        <Icon
          name="people"
          type="material"
          size={24}
          onPress={toggleShowMembers}
        />
      </View>
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
          {members.map((member) => (
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

      {/*members.map((member) => (
        <View>
          <Text>{`${member.surname} ${member.firstname} ${member.middlename}`}</Text>
          <Text>{member.roleType}</Text>
        </View>
      ))*/}
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
          {chatMessages &&
            chatMessages.map((msg, i) => (
              <View key={i} style={{ width: "100%" }}>
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
              </View>
            ))}
        </ScrollView>
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
            onPress={sendMessage}
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
