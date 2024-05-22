import { useAuth } from "@/providers/auth";
import { Faculties, columns } from "@/tables/faculties/columns";
import { FacultiesTable } from "@/tables/faculties/table";
import { useCallback, useEffect, useState } from "react";

export const DashboardPage = () => {
  const { token } = useAuth();

  const [faculties, setFaculties] = useState<Faculties[]>([]);

  const getFaculties = useCallback(async () => {
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/faculties", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      }).then(async (res) => {
        const json = await res.json();
        const data = [];
        for (let i = 0; i < json["faculties"].length; i++) {
          data.push({
            id: json["faculties"][i]["id"],
            name: json["faculties"][i]["name"],
          });
        }
        setFaculties(data);
      });
  }, [token]);

  useEffect(() => {
    getFaculties();
  }, [getFaculties]);

  return (
    <section>
      <FacultiesTable columns={columns} data={faculties} />
    </section>
  );
};
