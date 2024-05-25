import { Toaster } from "@/components/ui/toaster";
import MainLayout from "@/layouts/MainLayout";
import { AuthProvider, useAuth } from "@/providers/auth";
import { FC, useEffect } from "react";
import { Outlet } from "react-router-dom";

export const Root: FC = () => {
  const { refresh } = useAuth();
  useEffect(() => {
    refresh();
  }, [refresh]);
  return (
    <AuthProvider>
      <MainLayout>
        <Outlet />
      </MainLayout>
      <Toaster />
    </AuthProvider>
  );
};
