import Combobox from "@/components/Combobox";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
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
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { useApi } from "@/providers/api";
import { useAuth } from "@/providers/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { Check, ChevronsUpDown } from "lucide-react";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

interface EditDepartmentProps {
  id: string;
  name: string;
  faculty_id: string;
}

const formSchema = z.object({
  id: z.string().uuid(),
  name: z.string().optional(),
  faculty_id: z.string().uuid().optional(),
});

const EditDepartment: FC<EditDepartmentProps> = (props) => {
  const { id, name, faculty_id } = props;
  const { token } = useAuth();
  const { editDepartment, faculties } = useApi();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: id,
      name: name,
      faculty_id: faculty_id,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return editDepartment(token, values);
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
          <FormField
            control={form.control}
            name="faculty_id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Название кафедры</FormLabel>
                <FormControl>
                  <Combobox
                    value={field.value as string}
                    array={faculties}
                    placeholder="Выберите факультет"
                    notfound="Не найдено факультетов"
                    setValue={(value: string) =>
                      form.setValue("faculty_id", value)
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

export default EditDepartment;
