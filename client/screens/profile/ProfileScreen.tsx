import { SafeAreaView } from "react-native-safe-area-context";
import { Button, Chip, Dialog, Icon, Input, Text } from "@rneui/base";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { BottomTabScreenProps } from "@react-navigation/bottom-tabs";
import { MainTabsParamList } from "../main/MainBottomNavigator";
import { CompositeScreenProps } from "@react-navigation/native";
import { StackScreenProps } from "@react-navigation/stack";
import { AuthStackParamList } from "../auth/AuthStack";
import { useEffect, useState } from "react";
import { api, chatAPI } from "../../lib/config";
import { KeyboardAvoidingView, View } from "react-native";

type ProfileScreenProps = CompositeScreenProps<
  BottomTabScreenProps<MainTabsParamList, "Profile">,
  StackScreenProps<AuthStackParamList>
>;

interface Profile {
  description: string;
  fio: string;
  role: string;
  tags: {
    id: string;
    name: string;
  }[];
}

const ProfileScreen = ({ navigation }: ProfileScreenProps) => {
  const [profile, setProfile] = useState<Profile | null>(null);
  const [tag, setTag] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [showCreateTag, setShowCreateTag] = useState(false);
  const [showChangeDescription, setShowChangeDescription] = useState(false);

  const toggleShowCreateTag = () => {
    setShowCreateTag(!showCreateTag);
  };
  const toggleShowChangeDescription = () => {
    setShowChangeDescription(!showChangeDescription);
  };

  const logout = async () => {
    try {
      await AsyncStorage.clear();
      navigation.navigate("Login");
    } catch (e) { }
  };

  const changeProfile = async () => {
    const token = await AsyncStorage.getItem("access_token");

    const res = await fetch(chatAPI + "/api/v1/chat/change-profile", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        tags: profile?.tags ? [...profile.tags, tag] : [tag],
        description: description,
      }),
    });
    const json = await res.json();
    setProfile(json);
    setDescription(json["description"]);
    console.log(json);
  };

  const getProfile = async () => {
    const token = await AsyncStorage.getItem("access_token");

    const res = await fetch(chatAPI + "/api/v1/chat/get-profile", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    });
    const json = await res.json();
    setProfile(json);
    setDescription(json["description"]);
    console.log(json);
  };

  useEffect(() => {
    getProfile();
  }, []);

  return (
    <SafeAreaView>
      <View
        style={{
          marginTop: 64,
          paddingHorizontal: 16,
        }}
      >
        <Text
          style={{
            fontSize: 20,
          }}
        >
          {profile?.fio}
        </Text>
        <Text
          style={{
            fontSize: 14,
            marginTop: 8,
          }}
        >
          {profile?.description}
        </Text>

        {profile?.tags && profile?.tags.length > 0 ? (
          <>
            <Text
              style={{
                marginTop: 12,
              }}
            >
              Ваши теги:
            </Text>

            <View
              style={{
                flexDirection: "row",
                flexWrap: "wrap",
                gap: 8,
                marginTop: 4,
              }}
            >
              {profile.tags.map((profileTag) => (
                <Chip
                  key={profileTag.id}
                  title={profileTag.name}
                  icon={{
                    name: "close",
                    type: "font-awesome",
                    size: 20,
                  }}
                  onPress={() => console.log("Icon chip was pressed!")}
                  iconRight
                  type="outline"
                />
              ))}
            </View>
          </>
        ) : null}

        <Button
          containerStyle={{ marginVertical: 12 }}
          onPress={toggleShowCreateTag}
        >
          Добавить тег
        </Button>

        <Button
          containerStyle={{ marginBottom: 12 }}
          onPress={toggleShowChangeDescription}
        >
          Изменить описание
        </Button>
        <Dialog
          isVisible={showChangeDescription}
          onBackdropPress={toggleShowChangeDescription}
          overlayStyle={{
            backgroundColor: "#fff",
            marginBottom: 333,
          }}
          statusBarTranslucent
        >
          <Dialog.Title title="Изменить описание" />
          <KeyboardAvoidingView
            style={{
              flexDirection: "row",
              marginTop: 12,
              alignItems: "flex-start",
            }}
          >
            <View
              style={{
                flex: 1,
              }}
            >
              <Input
                value={description}
                multiline={true}
                onChangeText={(value) => setDescription(value)}
              />
            </View>
            <Icon
              onPress={() => console.log("hey")}
              containerStyle={{
                backgroundColor: "#2089dc",
                padding: 8,
                borderRadius: 12,
              }}
              name="send"
              color="#fff"
              size={20}
            />
          </KeyboardAvoidingView>
        </Dialog>
        <Dialog
          isVisible={showCreateTag}
          onBackdropPress={toggleShowCreateTag}
          overlayStyle={{
            backgroundColor: "#fff",
            marginBottom: 200,
          }}
          statusBarTranslucent
        >
          <Dialog.Title title="Создание тега" />
          <View
            style={{
              flexDirection: "row",
              marginTop: 12,
              alignItems: "flex-start",
            }}
          >
            <View
              style={{
                flex: 1,
              }}
            >
              <Input value={tag} onChangeText={(value) => setTag(value)} />
            </View>
            <Icon
              onPress={() => console.log("hey")}
              containerStyle={{
                backgroundColor: "#2089dc",
                padding: 8,
                borderRadius: 12,
              }}
              name="send"
              color="#fff"
              size={20}
            />
          </View>
        </Dialog>
        <Button onPress={logout}>Выход</Button>
      </View>
    </SafeAreaView>
  );
};

export default ProfileScreen;
