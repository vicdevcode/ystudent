import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogHeader,
  DialogContent,
  DialogTitle,
  DialogTrigger,
  DialogFooter,
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
import { useToast } from "@/components/ui/use-toast";
import { useAuth } from "@/providers/auth";
import { Faculties, columns } from "@/tables/faculties/columns";
import { FacultiesTable } from "@/tables/faculties/table";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

export const FacultiesPage = () => {
  const { token } = useAuth();

  const [faculties, setFaculties] = useState<Faculties[]>([]);

  const getFaculties = useCallback(async () => {
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/faculties", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      }).then(async (res) => {
        const json = await res.json();
        if (res.status === 401) console.log("hey");
        if (res.status === 200) {
          const data = [];
          for (let i = 0; i < json["faculties"].length; i++) {
            data.push({
              id: json["faculties"][i]["id"],
              name: json["faculties"][i]["name"],
            });
          }
          setFaculties(data);
        }
      });
  }, [token]);

  useEffect(() => {
    getFaculties();
  }, [getFaculties]);

  return (
    <section>
      <FacultiesTable columns={columns} data={faculties} />
      <AddFaculty />
    </section>
  );
};

const formSchema = z.object({
  name: z.string(),
});

const AddFaculty = () => {
  const { token, refresh } = useAuth();
  const { toast } = useToast();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/faculty/", {
        method: "POST",
        headers: {
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify(values),
      }).then(async (res) => {
        const json = await res.json();
        if (res.status === 401) refresh();
        if (res.status === 200) {
          toast({
            title: "OK",
            content: "Факультет был создан",
          });
        } else {
          toast({
            variant: "destructive",
            title: "ERROR",
            content: json["message"],
          });
        }
      });
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Добавить</Button>
      </DialogTrigger>
      <DialogContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <DialogHeader>
              <DialogTitle>Создание факультета</DialogTitle>
            </DialogHeader>
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Название факультета</FormLabel>
                  <FormControl>
                    <Input placeholder="Название" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <DialogFooter>
              <Button type="submit">Создать</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
