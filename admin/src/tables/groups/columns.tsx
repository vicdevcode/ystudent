import { ColumnDef } from "@tanstack/react-table";

export type Groups = {
  id: string;
  name: string;
  department_name: string;
};

export const columns: ColumnDef<Groups>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "name",
    header: "Название",
  },
  {
    accessorKey: "department_name",
    header: "Кафедра",
  },
];
