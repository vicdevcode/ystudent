import { useApi } from "@/providers/api";
import { columns } from "@/tables/groups/columns";
import { GroupsTable } from "@/tables/groups/table";
import { FC } from "react";

const GroupsPage: FC = () => {
  const { groups } = useApi();

  return (
    <section>
      <GroupsTable columns={columns} data={groups} />
    </section>
  );
};

export default GroupsPage;
