import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
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
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { z } from "zod";

const formSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});

export const AuthPage = () => {
  const { isAuth, setToken, setAuth } = useAuth();
  const { toast } = useToast();
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    return fetch(import.meta.env.VITE_AUTH_API + "/", {
      method: "POST",
      body: JSON.stringify(values),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status === 200) {
        localStorage.setItem("token", json["access_token"]);
        setToken(json["access_token"]);
        setAuth(true);
        navigate("/");
      } else {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        console.error(json["message"]);
      }
    });
  };

  useEffect(() => {
    if (isAuth) navigate("/");
  }, [navigate, isAuth]);

  return (
    <section className="h-screen flex justify-center items-center">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <Card className="w-[400px]">
            <CardHeader>
              <CardTitle>YStudent</CardTitle>
              <CardDescription>
                Авторизуйтесь, чтобы попасть в административную панель YStudent
              </CardDescription>
            </CardHeader>
            <CardContent>
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Электронная почта</FormLabel>
                    <FormControl>
                      <Input placeholder="e-mail" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Пароль</FormLabel>
                    <FormControl>
                      <Input
                        type="password"
                        placeholder="password"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </CardContent>
            <CardFooter>
              <Button type="submit">Авторизоваться</Button>
            </CardFooter>
          </Card>
        </form>
      </Form>
    </section>
  );
};
