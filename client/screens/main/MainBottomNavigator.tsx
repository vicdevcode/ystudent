import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import ProfileScreen from "../profile/ProfileScreen";
import SearchScreen from "../search/SearchScreen";
import ChatListScreen from "../chat/ChatListScreen";
import { Icon } from "@rneui/base";

export type MainTabsParamList = {
  Page: undefined;
  Chat: undefined;
  Search: undefined;
  Profile: undefined;
};

const Tab = createBottomTabNavigator<MainTabsParamList>();

const MainTabs = () => {
  return (
    <Tab.Navigator
      screenOptions={({ route }) => ({
        headerShown: false,
        tabBarIcon: ({ focused, color, size }) => {
          let iconName = "";

          if (route.name === "Chat") {
            iconName = focused ? "chat" : "chat-bubble-outline";
          } else if (route.name === "Search") {
            iconName = focused ? "search" : "search";
          } else if (route.name === "Profile") {
            iconName = focused ? "account-circle" : "account-circle";
          }

          return (
            <Icon name={iconName} type="material" size={size} color={color} />
          );
        },
      })}
    >
      <Tab.Screen name="Chat" component={ChatListScreen} />
      <Tab.Screen name="Search" component={SearchScreen} />
      <Tab.Screen name="Profile" component={ProfileScreen} />
    </Tab.Navigator>
  );
};

export default MainTabs;
