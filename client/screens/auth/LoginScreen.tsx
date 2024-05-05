import { Button, Input } from "@rneui/base";
import { SafeAreaView } from "react-native-safe-area-context";
import type { NativeStackScreenProps } from "@react-navigation/native-stack";
import { AuthStackParamList } from "./AuthStack";
import { useState } from "react";
import { authAPI } from "../../lib/config";
import { setSocket } from "../../lib/socket";

type LoginScreenProps = NativeStackScreenProps<AuthStackParamList, "Login">;

const LoginScreen = ({ navigation }: LoginScreenProps) => {
  const [email, onChangeEmail] = useState("");
  const [password, onChangePassword] = useState("");

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
    if (response.status != 200) return;
    const json = await response.json();

    setSocket(json["access_token"]);
    navigation.push("Main");
  };

  return (
    <SafeAreaView>
      <Input value={email} onChangeText={onChangeEmail} placeholder="Email" />
      <Input
        value={password}
        onChangeText={onChangePassword}
        secureTextEntry={true}
        placeholder="Password"
      />

      <Button onPress={login}>Войти</Button>
    </SafeAreaView>
  );
};

export default LoginScreen;
