import { useApi } from "@/providers/api";
import { columns } from "@/tables/faculties/columns";
import { FacultiesTable } from "@/tables/faculties/table";

const FacultiesPage = () => {
  const { faculties } = useApi();

  return (
    <section>
      <FacultiesTable columns={columns} data={faculties} />
    </section>
  );
};

export default FacultiesPage;
