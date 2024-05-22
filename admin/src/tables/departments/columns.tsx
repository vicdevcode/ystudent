import { ColumnDef } from "@tanstack/react-table";

export type Departments = {
  id: string;
  name: string;
  faculty_name: string;
};

export const columns: ColumnDef<Departments>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "name",
    header: "Название",
  },
  {
    accessorKey: "faculty_name",
    header: "Факультет",
  },
];
