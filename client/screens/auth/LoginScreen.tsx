import { Button, Input, Overlay, Text } from "@rneui/base";
import { SafeAreaView } from "react-native-safe-area-context";
import type { NativeStackScreenProps } from "@react-navigation/native-stack";
import { AuthStackParamList } from "./AuthStack";
import { useEffect, useState } from "react";
import { authAPI } from "../../lib/config";
import { setSocket } from "../../lib/socket";
import AsyncStorage from "@react-native-async-storage/async-storage";

type LoginScreenProps = NativeStackScreenProps<AuthStackParamList, "Login">;

const LoginScreen = ({ navigation }: LoginScreenProps) => {
  const [email, onChangeEmail] = useState("");
  const [password, onChangePassword] = useState("");
  const [errorVisible, setErrorVisible] = useState(false);
  const [errorContent, setErrorContent] = useState("");

  const checkAuth = async () => {
    const value = await AsyncStorage.getItem("access_token");
    if (value !== null) {
      setSocket(value);
      navigation.navigate("Main");
    }
  };

  const login = async () => {
    const response = await fetch(authAPI + "/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    });
    const json = await response.json();
    if (response.status != 200) {
      setErrorVisible(true);
      setErrorContent(json["message"]);
      return;
    }

    setSocket(json["access_token"]);
    navigation.push("Main");
  };

  const understandableButton = () => setErrorVisible(false);

  useEffect(() => {
    checkAuth();
  }, []);

  return (
    <SafeAreaView>
      <Input value={email} onChangeText={onChangeEmail} placeholder="Email" />
      <Input
        value={password}
        onChangeText={onChangePassword}
        secureTextEntry={true}
        placeholder="Password"
      />
      <Overlay isVisible={errorVisible} onBackdropPress={understandableButton}>
        <Text h4>ОШИБКА</Text>
        <Text>{errorContent}</Text>
        <Button onPress={understandableButton}>Понятно</Button>
      </Overlay>

      <Button onPress={login}>Войти</Button>
    </SafeAreaView>
  );
};

export default LoginScreen;
