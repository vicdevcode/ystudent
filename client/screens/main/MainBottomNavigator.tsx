import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import MainScreen from "./MainScreen";
import ProfileScreen from "../profile/ProfileScreen";
import ChatScreen from "../chat/ChatScreen";
import SearchScreen from "../search/SearchScreen";

export type MainTabsParamList = {
  Page: undefined;
  Chat: undefined;
  Search: undefined;
  Profile: undefined;
};

const Tab = createBottomTabNavigator<MainTabsParamList>();

const MainTabs = () => {
  return (
    <Tab.Navigator>
      <Tab.Screen name="Page" component={MainScreen} />
      <Tab.Screen name="Chat" component={ChatScreen} />
      <Tab.Screen name="Search" component={SearchScreen} />
      <Tab.Screen name="Profile" component={ProfileScreen} />
    </Tab.Navigator>
  );
};

export default MainTabs;
