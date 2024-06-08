import { ColumnDef } from "@tanstack/react-table";
import CreateDepartment from "./create";
import { ActionButtonDepartment } from "./action";

export type Departments = {
  id: string;
  name: string;
  faculty_name: string;
  faculty_id: string;
};

export const columns: ColumnDef<Departments>[] = [
  {
    accessorKey: "id",
    header: "id",
    size: 320,
  },
  {
    accessorKey: "name",
    header: "Название",
  },
  {
    accessorKey: "faculty_id",
    header: "",
  },
  {
    accessorKey: "faculty_name",
    header: "Факультет",
  },
  {
    id: "actions",
    enableHiding: false,
    size: 50,
    enableResizing: false,
    header: () => <CreateDepartment />,
    cell: ({ row }) => (
      <ActionButtonDepartment
        id={row.getValue("id")}
        name={row.getValue("name")}
        faculty_id={row.getValue("faculty_id")}
      />
    ),
  },
];
