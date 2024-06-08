import AlertDialogDelete from "@/components/AlertDialogDelete";
import { useApi } from "@/providers/api";
import { FC } from "react";

interface DeleteProps {
  id: string;
}

const DeleteEmployee: FC<DeleteProps> = ({ id }) => {
  const { deleteEmployee } = useApi();

  return (
    <AlertDialogDelete
      id={id}
      description="Вы навсегда удалите этого сотрудника"
      onClick={deleteEmployee}
    />
  );
};

export default DeleteEmployee;
