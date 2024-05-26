import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Root } from "./pages/Root";
import { DashboardPage } from "./pages/DashboardPage";
import { AuthPage } from "./pages/AuthPage";
import FacultiesPage from "./pages/FacultiesPage";
import DepartmentPage from "./pages/DepartmentsPage";
import GroupsPage from "./pages/GroupsPage";
import TeachersPage from "./pages/TeachersPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "",
        element: <DashboardPage />,
      },
      {
        path: "faculties",
        element: <FacultiesPage />,
      },
      {
        path: "departments",
        element: <DepartmentPage />,
      },
      {
        path: "groups",
        element: <GroupsPage />,
      },
      {
        path: "teachers",
        element: <TeachersPage />,
      },
      {
        path: "auth",
        element: <AuthPage />,
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
}
