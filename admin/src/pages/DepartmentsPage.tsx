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
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
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
import { useToast } from "@/components/ui/use-toast";
import { cn } from "@/lib/utils";
import { useAuth } from "@/providers/auth";
import { Departments, columns } from "@/tables/departments/columns";
import { DepartmentsTable } from "@/tables/departments/table";
import { Faculties } from "@/tables/faculties/columns";
import { zodResolver } from "@hookform/resolvers/zod";
import { Check, ChevronsUpDown } from "lucide-react";
import { useCallback, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const DepartmentPage = () => {
  const { token } = useAuth();

  const [departments, setDepartments] = useState<Departments[]>([]);

  const getDepartments = useCallback(async () => {
    console.log(token);
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/departments", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      }).then(async (res) => {
        console.log(res.status);
        const json = await res.json();
        if (res.status === 200) {
          const data = [];
          for (let i = 0; i < json["departments"].length; i++) {
            data.push({
              id: json["departments"][i]["id"],
              name: json["departments"][i]["name"],
              faculty_name: json["departments"][i]["faculty_name"],
            });
          }
          setDepartments(data);
        }
      });
  }, [token]);

  useEffect(() => {
    getDepartments();
  }, [getDepartments]);

  return (
    <section>
      <DepartmentsTable columns={columns} data={departments} />
      <AddDepartment />
    </section>
  );
};

const formSchema = z.object({
  name: z.string(),
  faculty_id: z.string(),
});

const AddDepartment = () => {
  const { token, refresh } = useAuth();
  const { toast } = useToast();

  const [faculties, setFaculties] = useState<Faculties[]>([
    {
      id: "",
      name: "",
    },
  ]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    if (token)
      return fetch(import.meta.env.VITE_MAIN_API + "/department/", {
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
            content: "Кафедра была создана",
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
    <Dialog>
      <DialogTrigger asChild>
        <Button>Добавить</Button>
      </DialogTrigger>
      <DialogContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <DialogHeader>
              <DialogTitle>Создание кафедры</DialogTitle>
            </DialogHeader>
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Название кафедры</FormLabel>
                  <FormControl>
                    <Input placeholder="Название" {...field} />
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
                  <FormLabel>Факультет</FormLabel>
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant="outline"
                          role="combobox"
                          className={cn(
                            "w-[200px] justify-between",
                            !field.value && "text-muted-foreground",
                          )}
                        >
                          {field.value
                            ? faculties.find(
                              (language) => language.id === field.value,
                            )?.name
                            : "Выберите факультет"}
                          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                      <Command>
                        <CommandInput placeholder="Выберите факультет" />
                        <CommandList>
                          <CommandEmpty>Не найдено факультетов</CommandEmpty>
                          <CommandGroup>
                            {faculties.map((faculty) => (
                              <CommandItem
                                value={faculty.name}
                                key={faculty.id}
                                onSelect={() => {
                                  form.setValue("faculty_id", faculty.id);
                                }}
                              >
                                <Check
                                  className={cn(
                                    "mr-2 h-4 w-4",
                                    faculty.id === field.value
                                      ? "opacity-100"
                                      : "opacity-0",
                                  )}
                                />
                                {faculty.name}
                              </CommandItem>
                            ))}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
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

export default DepartmentPage;
