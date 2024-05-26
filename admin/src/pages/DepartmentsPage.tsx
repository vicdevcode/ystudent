import { useApi } from "@/providers/api";
import { columns } from "@/tables/departments/columns";
import { DepartmentsTable } from "@/tables/departments/table";

const DepartmentPage = () => {
  const { departments } = useApi();

  return (
    <section>
      <DepartmentsTable columns={columns} data={departments} />
    </section>
  );
};

export default DepartmentPage;
