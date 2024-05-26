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

interface DeleteFacultyProps {
  id: string;
}

const DeleteFaculty: FC<DeleteFacultyProps> = ({ id }) => {
  const { token } = useAuth();
  const { deleteFaculty } = useApi();

  return (
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Вы уверены?</AlertDialogTitle>
        <AlertDialogDescription>
          Вы навсегда удалите этот факультет
        </AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Отмена</AlertDialogCancel>
        <AlertDialogAction onClick={async () => deleteFaculty(token, id)}>
          Удалить
        </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  );
};

export default DeleteFaculty;
