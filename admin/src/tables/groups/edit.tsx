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

interface EditGroupProps {
  id: string;
  name: string;
  department_id: string;
  curator_id: string;
}

const formSchema = z.object({
  id: z.string().uuid(),
  name: z.string().optional(),
  department_id: z.string().uuid().optional(),
  curator_id: z.string().uuid().optional(),
});

const EditGroup: FC<EditGroupProps> = (props) => {
  const { id, name, department_id, curator_id } = props;
  const { token } = useAuth();
  const { editGroup, departments, teachers } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: id,
      name: name,
      department_id: department_id,
      curator_id: curator_id,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return editGroup(token, values);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Редактирование группы</DialogTitle>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Название группы</FormLabel>
                <FormControl>
                  <Input placeholder="Новое название группы" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="department_id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Кафедра</FormLabel>
                <FormControl>
                  <Combobox
                    value={field.value as string}
                    array={departments}
                    placeholder="Выберите кафедру"
                    notfound="Не найдено кафедр"
                    setValue={(value: string) =>
                      form.setValue("department_id", value)
                    }
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="curator_id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Куратор</FormLabel>
                <FormControl>
                  <Combobox
                    value={field.value as string}
                    array={teachers}
                    placeholder="Выберите куратора"
                    notfound="Не найдено преподавателей"
                    setValue={(value: string) =>
                      form.setValue("curator_id", value)
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

export default EditGroup;
