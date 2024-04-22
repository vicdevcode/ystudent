import { Button, Input } from "@rneui/base";
import { SafeAreaView } from "react-native-safe-area-context";
import type { NativeStackScreenProps } from "@react-navigation/native-stack";
import { AuthStackParamList } from "./AuthStack";

type LoginScreenProps = NativeStackScreenProps<AuthStackParamList, "Login">;

const LoginScreen = ({ navigation }: LoginScreenProps) => {
  return (
    <SafeAreaView>
      <Input placeholder="Email" />
      <Input placeholder="Password" />

      <Button onPress={() => navigation.navigate("Main")}>Войти</Button>
    </SafeAreaView>
  );
};

export default LoginScreen;
