import Combobox from "@/components/Combobox";
import { Button } from "@/components/ui/button";
import {
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";
import { useApi } from "@/providers/api";
import { useAuth } from "@/providers/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

interface AddAdminFacultyProps {
  id: string;
}

const formSchema = z.object({
  id: z.string().uuid(),
  user_id: z.string().uuid(),
  type: z.enum(["faculty", "department"]),
});

const AddAdminFaculty: FC<AddAdminFacultyProps> = (props) => {
  const { id } = props;
  const { token } = useAuth();
  const { addChatAdmin, employees } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: id,
      user_id: "",
      type: "faculty",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return addChatAdmin(token, values);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Добавить админа</DialogTitle>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="user_id"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <Combobox
                    value={field.value as string}
                    array={employees}
                    is_user_id
                    placeholder="Выберите редактора"
                    notfound="Не найдено сотрудников"
                    setValue={(value: string) =>
                      form.setValue("user_id", value)
                    }
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <DialogFooter>
            <Button type="submit">Добавить</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default AddAdminFaculty;
