import { SafeAreaView } from "react-native-safe-area-context";
import { Button } from "@rneui/base";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { BottomTabScreenProps } from "@react-navigation/bottom-tabs";
import { MainTabsParamList } from "../main/MainBottomNavigator";
import { CompositeScreenProps } from "@react-navigation/native";
import { StackScreenProps } from "@react-navigation/stack";
import { AuthStackParamList } from "../auth/AuthStack";

type ProfileScreenProps = CompositeScreenProps<
  BottomTabScreenProps<MainTabsParamList, "Profile">,
  StackScreenProps<AuthStackParamList>
>;

const ProfileScreen = ({ navigation }: ProfileScreenProps) => {
  const logout = async () => {
    try {
      await AsyncStorage.clear();
      navigation.navigate("Login");
    } catch (e) { }
  };

  return (
    <SafeAreaView>
      <Button onPress={logout}>Выход</Button>
    </SafeAreaView>
  );
};

export default ProfileScreen;
