import AsyncStorage from "@react-native-async-storage/async-storage";
import { Button, Chip, SearchBar, Text } from "@rneui/base";
import { useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { chatAPI } from "../../lib/config";
import { View } from "react-native";
import { ScrollView } from "react-native-gesture-handler";
import { roleTypes } from "../../lib/utils";

interface Profile {
  id: string;
  description: string;
  fio: string;
  role: string;
  tags: {
    id: string;
    name: string;
  }[];
}
const SearchScreen = () => {
  const [search, setSearch] = useState("");
  const [result, setResult] = useState<Profile[]>([]);

  const updateSearch = (search: string) => {
    setSearch(search);
    getProfiles(search);
  };

  const getProfiles = async (tag: string) => {
    const token = await AsyncStorage.getItem("access_token");

    const res = await fetch(chatAPI + "/api/v1/chat/search/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        tag: tag,
      }),
    });
    const json = await res.json();
    setResult(json);
    console.log(json);
  };

  return (
    <SafeAreaView>
      <SearchBar
        lightTheme
        placeholder="Поиск"
        onChangeText={updateSearch}
        value={search}
      />
      <ScrollView
        style={{
          paddingHorizontal: 8,
          marginTop: 12,
        }}
      >
        {result.length > 0 ? (
          result.map((profile) => (
            <View
              key={profile.id}
              style={{
                padding: 8,
                backgroundColor: "#e2e2e2",
                borderRadius: 12,
              }}
            >
              <Text style={{ fontSize: 16 }}>{profile.fio}</Text>
              <Text style={{ fontSize: 12, color: "#999" }}>
                {profile.role === "STUDENT" ||
                  (profile.role === "TEACHER" && roleTypes[profile.role])}
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
                    size="sm"
                    titleStyle={{ fontSize: 12 }}
                    title={profileTag.name}
                  />
                ))}
              </View>
              <Button
                title="Написать"
                buttonStyle={{ borderRadius: 12 }}
                containerStyle={{ marginTop: 8 }}
              />
            </View>
          ))
        ) : (
          <Text>Не найдено</Text>
        )}
      </ScrollView>
    </SafeAreaView>
  );
};

export default SearchScreen;
