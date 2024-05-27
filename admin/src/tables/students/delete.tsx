import AlertDialogDelete from "@/components/AlertDialogDelete";
import { useApi } from "@/providers/api";
import { FC } from "react";

interface DeleteProps {
  id: string;
}

const DeleteStudent: FC<DeleteProps> = ({ id }) => {
  const { deleteStudent } = useApi();

  return (
    <AlertDialogDelete
      id={id}
      description="Вы навсегда удалите этого студента"
      onClick={deleteStudent}
    />
  );
};

export default DeleteStudent;
