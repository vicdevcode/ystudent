import {
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { useApi } from "@/providers/api";
import { useAuth } from "@/providers/auth";
import { FC } from "react";

interface DeleteDepartmentProps {
  id: string;
}

const DeleteDepartment: FC<DeleteDepartmentProps> = ({ id }) => {
  const { token } = useAuth();
  const { deleteDepartment } = useApi();

  return (
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Вы уверены?</AlertDialogTitle>
        <AlertDialogDescription>
          Вы навсегда удалите эту кафедру
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Отмена</AlertDialogCancel>
        <AlertDialogAction onClick={async () => deleteDepartment(token, id)}>
          Удалить
        </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  );
};

export default DeleteDepartment;
