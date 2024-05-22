import { Toaster } from "@/components/ui/toaster";
import MainLayout from "@/layouts/MainLayout";
import { AuthProvider } from "@/providers/auth";
import { FC } from "react";
import { Outlet } from "react-router-dom";

export const Root: FC = () => {
  return (
    <AuthProvider>
      <MainLayout>
        <Outlet />
      </MainLayout>
      <Toaster />
    </AuthProvider>
  );
};
