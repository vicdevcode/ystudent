import { Button } from "@/components/ui/button";
import {
  DialogContent,
  DialogDescription,
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

interface EditFacultyProps {
  id: string;
  name: string;
}

const formSchema = z.object({
  id: z.string().uuid(),
  name: z.string(),
});

const EditFaculty: FC<EditFacultyProps> = (props) => {
  const { id, name } = props;
  const { token } = useAuth();
  const { editFaculty } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: id,
      name: name,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return editFaculty(token, values);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Факультет</DialogTitle>
        <DialogDescription>
          Здесь можно отредактировать факультет
        </DialogDescription>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Название факультета</FormLabel>
                <FormControl>
                  <Input placeholder="Новое название факультета" {...field} />
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

export default EditFaculty;
