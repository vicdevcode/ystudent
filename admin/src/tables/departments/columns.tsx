import { ColumnDef } from "@tanstack/react-table";
import CreateDepartment from "./create";
import { AlertDialog, AlertDialogTrigger } from "@/components/ui/alert-dialog";
import { Dialog, DialogTrigger } from "@/components/ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal } from "lucide-react";
import EditDepartment from "./edit";
import DeleteDepartment from "./delete";

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
      <AlertDialog>
        <Dialog>
          <DropdownMenu>
            <DropdownMenuTrigger>
              <MoreHorizontal />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>
                <DialogTrigger>Редактировать</DialogTrigger>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem>
                <AlertDialogTrigger>Удалить</AlertDialogTrigger>
              </DropdownMenuItem>
            </DropdownMenuContent>
            <EditDepartment
              id={row.getValue("id")}
              name={row.getValue("name")}
              faculty_id={row.getValue("faculty_id")}
            />
            <DeleteDepartment id={row.getValue("id")} />
          </DropdownMenu>
        </Dialog>
      </AlertDialog>
    ),
  },
];
