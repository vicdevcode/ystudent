import { ColumnDef } from "@tanstack/react-table";
import CreateTeacher from "./create";
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
import EditTeacher from "./edit";
import DeleteTeacher from "./delete";

export type Employees = {
  id: string;
  user_id: string;
  firstname: string;
  middlename: string;
  surname: string;
  name: string;
  email: string;
};

export const columns: ColumnDef<Employees>[] = [
  {
    accessorKey: "id",
    header: "id",
    size: 300,
  },
  {
    accessorKey: "name",
    header: "ФИО",
  },
  {
    accessorKey: "firstname",
  },
  {
    accessorKey: "middlename",
  },
  {
    accessorKey: "surname",
  },
  {
    accessorKey: "user_id",
  },
  {
    accessorKey: "email",
    header: "Электронная почта",
  },
  {
    id: "actions",
    enableHiding: false,
    header: () => <CreateTeacher />,
    size: 50,
    cell: ({ row }) => {
      return (
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
              <EditTeacher
                id={row.getValue("id")}
                firstname={row.getValue("firstname")}
                middlename={row.getValue("middlename")}
                surname={row.getValue("surname")}
                email={row.getValue("email")}
              />
              <DeleteTeacher id={row.getValue("id")} />
            </DropdownMenu>
          </Dialog>
        </AlertDialog>
      );
    },
  },
];
