import Header from "@/components/Header";
import { FC, PropsWithChildren } from "react";
import { useLocation } from "react-router-dom";

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
  const location = useLocation();

  if (location.pathname === "/auth") return <main>{children}</main>;

  return (
    <main>
      <Header />
      <div className="container mx-auto my-0">{children}</div>
    </main>
  );
};

export default MainLayout;
