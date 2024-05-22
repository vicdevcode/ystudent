import { ColumnDef } from "@tanstack/react-table";

export type Faculties = {
  id: string;
  name: string;
};

export const columns: ColumnDef<Faculties>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "name",
    header: "Название",
  },
];
