import { ColumnDef } from "@tanstack/react-table";

export type Teachers = {
  id: string;
  name: string;
  email: string;
};

export const columns: ColumnDef<Teachers>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "name",
    header: "ФИО",
  },
  {
    accessorKey: "email",
    header: "Электронная почта",
  },
];
