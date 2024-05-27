import AlertDialogDelete from "@/components/AlertDialogDelete";
import { useApi } from "@/providers/api";
import { FC } from "react";

interface DeleteProps {
  id: string;
}

const DeleteTeacher: FC<DeleteProps> = ({ id }) => {
  const { deleteTeacher } = useApi();

  return (
    <AlertDialogDelete
      id={id}
      description="Вы навсегда удалите этого преподавателя"
      onClick={deleteTeacher}
    />
  );
};

export default DeleteTeacher;
