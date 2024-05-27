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
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useApi } from "@/providers/api";
import { useAuth } from "@/providers/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

interface EditProps {
  id: string;
  firstname: string;
  middlename: string;
  surname: string;
  email: string;
  group_id: string;
}

const formSchema = z.object({
  id: z.string(),
  firstname: z.string(),
  middlename: z.string().optional(),
  surname: z.string(),
  email: z.string(),
  group_id: z.string().uuid(),
});

const EditStudent: FC<EditProps> = (props) => {
  const { token } = useAuth();
  const { editStudent, groups } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: props.id,
      firstname: props.firstname,
      middlename: props.middlename,
      surname: props.surname,
      email: props.email,
      group_id: props.group_id,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return editStudent(token, values);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Редактирование студента</DialogTitle>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="surname"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Фамилия</FormLabel>
                <FormControl>
                  <Input placeholder="Фамилия" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="firstname"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Имя</FormLabel>
                <FormControl>
                  <Input placeholder="Имя" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="middlename"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Отчетство</FormLabel>
                <FormControl>
                  <Input placeholder="Отчетство" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Электронная почта</FormLabel>
                <FormControl>
                  <Input placeholder="email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="group_id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Группа</FormLabel>
                <FormControl>
                  <Combobox
                    value={field.value as string}
                    array={groups}
                    placeholder="Выберите группу"
                    notfound="Не найдено групп"
                    setValue={(value: string) =>
                      form.setValue("group_id", value)
                    }
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <DialogFooter>
            <Button type="submit">Отредактировать</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default EditStudent;
