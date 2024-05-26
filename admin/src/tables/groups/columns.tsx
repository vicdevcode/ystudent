import { AlertDialog, AlertDialogTrigger } from "@/components/ui/alert-dialog";
import { Dialog, DialogTrigger } from "@/components/ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ColumnDef } from "@tanstack/react-table";
import { MoreHorizontal } from "lucide-react";
import EditGroup from "./edit";
import DeleteGroup from "./delete";
import CreateGroup from "./create";

export type Groups = {
  id: string;
  name: string;
  department_name: string;
  curator_fio: string;
};

export const columns: ColumnDef<Groups>[] = [
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
    accessorKey: "department_id",
  },
  {
    accessorKey: "curator_id",
  },
  {
    accessorKey: "department_name",
    header: "Кафедра",
  },
  {
    accessorKey: "curator_fio",
    header: "Куратор",
  },
  {
    id: "actions",
    enableHiding: false,
    header: () => <CreateGroup />,
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
              <EditGroup
                id={row.getValue("id")}
                name={row.getValue("name")}
                department_id={row.getValue("department_id")}
                curator_id={row.getValue("curator_id")}
              />
              <DeleteGroup id={row.getValue("id")} />
            </DropdownMenu>
          </Dialog>
        </AlertDialog>
      );
    },
  },
];
