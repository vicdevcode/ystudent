import { ColumnDef } from "@tanstack/react-table";
import CreateFaculty from "./create";
import { ActionButton } from "./action";

export type Faculties = {
  id: string;
  name: string;
};

export const columns: ColumnDef<Faculties>[] = [
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
    id: "actions",
    enableHiding: false,
    size: 50,
    enableResizing: false,
    header: () => <CreateFaculty />,
    cell: ({ row }) => (
      <ActionButton id={row.getValue("id")} name={row.getValue("name")} />
    ),
  },
];
