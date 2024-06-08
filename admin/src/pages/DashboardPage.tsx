import { useApi } from "@/providers/api";
import { columns as dColumns } from "@/tables/departments/columns";
import { DepartmentsTable } from "@/tables/departments/table";
import { columns as eColumns } from "@/tables/employees/columns";
import { EmployeesTable } from "@/tables/employees/table";
import { columns as fColumns } from "@/tables/faculties/columns";
import { FacultiesTable } from "@/tables/faculties/table";
import { columns as gColumns } from "@/tables/groups/columns";
import { GroupsTable } from "@/tables/groups/table";
import { columns as sColumns } from "@/tables/students/columns";
import { StudentsTable } from "@/tables/students/table";
import { columns as tColumns } from "@/tables/teachers/columns";
import { TeachersTable } from "@/tables/teachers/table";

export const DashboardPage = () => {
  const { faculties, departments, groups, students, teachers, employees } =
    useApi();

  return (
    <section>
      <h1 className="mb-4 text-xl font-semibold">Факультеты</h1>
      <FacultiesTable columns={fColumns} data={faculties} />
      <h1 className="mb-4 text-xl font-semibold">Кафедры</h1>
      <DepartmentsTable columns={dColumns} data={departments} />
      <h1 className="mb-4 text-xl font-semibold">Группы</h1>
      <GroupsTable columns={gColumns} data={groups} />
      <h1 className="mb-4 text-xl font-semibold">Пользователи</h1>
      <h2 className="mb-4 text-lg font-medium">Сотрудники</h2>
      <EmployeesTable columns={eColumns} data={employees} />
      <h2 className="mb-4 text-lg font-medium">Преподаватели</h2>
      <TeachersTable columns={tColumns} data={teachers} />
      <h2 className="mb-4 text-lg font-medium">Студенты</h2>
      <StudentsTable columns={sColumns} data={students} />
    </section>
  );
};
