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
}

const formSchema = z.object({
  id: z.string().uuid(),
  firstname: z.string().optional(),
  middlename: z.string().optional(),
  surname: z.string().optional(),
  email: z.string().optional(),
});

const EditEmployee: FC<EditProps> = ({
  id,
  firstname,
  middlename,
  surname,
  email,
}) => {
  const { token } = useAuth();
  const { editEmployee } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: id,
      firstname: firstname,
      middlename: middlename,
      surname: surname,
      email: email,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return editEmployee(token, values);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Редактирование сотрудника</DialogTitle>
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
          <DialogFooter>
            <Button type="submit">Отредактировать</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default EditEmployee;
