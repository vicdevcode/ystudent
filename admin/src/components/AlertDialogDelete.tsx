import {
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { useAuth } from "@/providers/auth";
import { FC } from "react";

interface AlertDialogDeleteProps {
  id: string;
  description: string;
  onClick: (token: string, id: string) => Promise<void>;
}

const AlertDialogDelete: FC<AlertDialogDeleteProps> = ({
  id,
  description,
  onClick,
}) => {
  const { token } = useAuth();

  return (
    <AlertDialogContent>
      <AlertDialogHeader>
        <AlertDialogTitle>Вы уверены?</AlertDialogTitle>
        <AlertDialogDescription>{description}</AlertDialogDescription>
      </AlertDialogHeader>
      <AlertDialogFooter>
        <AlertDialogCancel>Отмена</AlertDialogCancel>
        <AlertDialogAction onClick={async () => onClick(token, id)}>
          Удалить
        </AlertDialogAction>
      </AlertDialogFooter>
    </AlertDialogContent>
  );
};

export default AlertDialogDelete;
