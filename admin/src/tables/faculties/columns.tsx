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
import EditFaculty from "./edit";
import DeleteFaculty from "./delete";
import CreateFaculty from "./create";

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
            <EditFaculty id={row.getValue("id")} name={row.getValue("name")} />
            <DeleteFaculty id={row.getValue("id")} />
          </DropdownMenu>
        </Dialog>
      </AlertDialog>
    ),
  },
];
