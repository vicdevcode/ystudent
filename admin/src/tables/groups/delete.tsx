import AlertDialogDelete from "@/components/AlertDialogDelete";
import { useApi } from "@/providers/api";
import { FC } from "react";

interface DeleteProps {
  id: string;
}

const DeleteGroup: FC<DeleteProps> = ({ id }) => {
  const { deleteGroup } = useApi();

  return (
    <AlertDialogDelete
      id={id}
      description="Вы навсегда удалите эту группу"
      onClick={deleteGroup}
    />
  );
};

export default DeleteGroup;
