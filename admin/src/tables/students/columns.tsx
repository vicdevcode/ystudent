import { ColumnDef } from "@tanstack/react-table";
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
import CreateStudent from "./create";
import DeleteStudent from "./delete";
import EditStudent from "./edit";

export type Students = {
  id: string;
  firstname: string;
  middlename: string;
  surname: string;
  name: string;
  email: string;
  group_name: string;
  group_id: string;
};

export const columns: ColumnDef<Students>[] = [
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
    accessorKey: "group_id",
  },
  {
    accessorKey: "group_name",
    header: "Группа",
  },
  {
    accessorKey: "email",
    header: "Электронная почта",
  },
  {
    id: "actions",
    enableHiding: false,
    header: () => <CreateStudent />,
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
              <EditStudent
                id={row.getValue("id")}
                firstname={row.getValue("firstname")}
                middlename={row.getValue("middlename")}
                surname={row.getValue("surname")}
                email={row.getValue("email")}
                group_id={row.getValue("group_id")}
              />
              <DeleteStudent id={row.getValue("id")} />
            </DropdownMenu>
          </Dialog>
        </AlertDialog>
      );
    },
  },
];
