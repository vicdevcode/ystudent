import { createStackNavigator } from "@react-navigation/stack";
import LoginScreen from "./LoginScreen";
import MainScreen from "../main/MainScreen";

export type AuthStackParamList = {
  Login: undefined;
  Main: undefined;
};

const Stack = createStackNavigator<AuthStackParamList>();

const AuthStack = () => {
  return (
    <Stack.Navigator
      screenOptions={() => ({
        headerShown: false,
      })}
    >
      <Stack.Screen component={LoginScreen} name="Login" />
      <Stack.Screen component={MainScreen} name="Main" />
    </Stack.Navigator>
  );
};

export default AuthStack;
