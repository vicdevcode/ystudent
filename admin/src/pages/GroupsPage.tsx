import { useAuth } from "@/providers/auth";
import { columns } from "@/tables/faculties/columns";
import { Groups } from "@/tables/groups/columns";
import { GroupsTable } from "@/tables/groups/table";
import { FC, useCallback, useEffect, useState } from "react";

const GroupsPage: FC = () => {
  const { token } = useAuth();
  const [groups, setGroups] = useState<Groups[]>([]);

  const getGroups = useCallback(async () => {
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/groups", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      }).then(async (res) => {
        console.log(res.status);
        const json = await res.json();
        if (res.status === 200) {
          const data = [];
          for (let i = 0; i < json["groups"].length; i++) {
            data.push({
              id: json["groups"][i]["id"],
              name: json["groups"][i]["name"],
              department_name: json["groups"][i]["depatment_name"],
            });
          }
          setGroups(data);
        }
      });
  }, [token]);

  useEffect(() => {
    getGroups();
  }, [getGroups]);
  return (
    <section>
      <GroupsTable columns={columns} data={groups} />
    </section>
  );
};

export default GroupsPage;
