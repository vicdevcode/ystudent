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
import { FC, useState } from "react";
import EditDepartment from "./edit";
import DeleteDepartment from "./delete";
import AddAdminDepartment from "./add-admin";
import DeleteAdminDepartment from "./delete-admin";

interface ActionButtonProps {
  id: string;
  name: string;
  faculty_id: string;
}

export const ActionButtonDepartment: FC<ActionButtonProps> = (props) => {
  const { id, name, faculty_id } = props;

  const [dialogStatus, setDialogStatus] = useState("add");

  return (
    <AlertDialog>
      <Dialog>
        <DropdownMenu>
          <DropdownMenuTrigger>
            <MoreHorizontal />
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuItem>
              <DialogTrigger onClick={() => setDialogStatus("add admin")}>
                Добавить редактора
              </DialogTrigger>
            </DropdownMenuItem>
            <DropdownMenuItem>
              <DialogTrigger onClick={() => setDialogStatus("delete admin")}>
                Убрать редактора
              </DialogTrigger>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>
              <DialogTrigger onClick={() => setDialogStatus("edit")}>
                Редактировать
              </DialogTrigger>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>
              <AlertDialogTrigger>Удалить</AlertDialogTrigger>
            </DropdownMenuItem>
          </DropdownMenuContent>
          {dialogStatus === "add admin" && <AddAdminDepartment id={id} />}
          {dialogStatus === "delete admin" && <DeleteAdminDepartment id={id} />}
          {dialogStatus === "edit" && (
            <EditDepartment id={id} name={name} faculty_id={faculty_id} />
          )}
          <DeleteDepartment id={id} />
        </DropdownMenu>
      </Dialog>
    </AlertDialog>
  );
};
