import { useApi } from "@/providers/api";
import { columns } from "@/tables/teachers/columns";
import { TeachersTable } from "@/tables/teachers/table";
import { FC } from "react";

const TeachersPage: FC = () => {
  const { teachers } = useApi();

  return (
    <section>
      <TeachersTable columns={columns} data={teachers} />
    </section>
  );
};

export default TeachersPage;
