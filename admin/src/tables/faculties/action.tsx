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
import { FC, useState } from "react";
import { DialogContent } from "@radix-ui/react-dialog";
import AddAdminFaculty from "./add-admin";
import DeleteAdminFaculty from "./delete-admin";

interface ActionButtonProps {
  id: string;
  name: string;
}

export const ActionButton: FC<ActionButtonProps> = (props) => {
  const { id, name } = props;

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
          {dialogStatus === "add admin" && <AddAdminFaculty id={id} />}
          {dialogStatus === "delete admin" && <DeleteAdminFaculty id={id} />}
          {dialogStatus === "edit" && <EditFaculty id={id} name={name} />}

          <DeleteFaculty id={id} />
        </DropdownMenu>
      </Dialog>
    </AlertDialog>
  );
};
