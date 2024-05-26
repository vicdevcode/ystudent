import { Toaster } from "@/components/ui/toaster";
import MainLayout from "@/layouts/MainLayout";
import ApiProvider from "@/providers/api";
import { AuthProvider } from "@/providers/auth";
import { FC } from "react";
import { Outlet } from "react-router-dom";

export const Root: FC = () => {
  return (
    <AuthProvider>
      <ApiProvider>
        <MainLayout>
          <Outlet />
        </MainLayout>
      </ApiProvider>
      <Toaster />
    </AuthProvider>
  );
};
