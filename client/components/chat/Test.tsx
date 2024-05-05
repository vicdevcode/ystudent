import { useEffect, useState } from "react";
import { Text, View } from "react-native";
import { socketIO } from "../../lib/socket";
import { chatAPI } from "../../lib/config";

const TestChat = () => {
  const [hasConnection, setConnection] = useState(false);
  const [time, setTime] = useState<string | null>(null);

  useEffect(function didMount() {
    if (socketIO.socket) {
      socketIO.socket.io.on("open", () => setConnection(true));
      socketIO.socket.io.on("close", () => setConnection(false));

      socketIO.socket.on("time-msg", (data) => {
        setTime(new Date(data.time).toString());
      });
    }
    return function didUnmount() {
      if (socketIO.socket) {
        socketIO.socket.disconnect();
        socketIO.socket.removeAllListeners();
      }
    };
  }, []);

  return (
    <View>
      {!hasConnection && (
        <>
          <Text>Connecting to {chatAPI}...</Text>
          <Text>Make sure the backend is started and reachable</Text>
        </>
      )}

      {hasConnection && (
        <>
          <Text>Server time</Text>
          <Text>{time}</Text>
        </>
      )}
    </View>
  );
};

export default TestChat;
