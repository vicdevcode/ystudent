import { useApi } from "@/providers/api";
import { columns } from "@/tables/employees/columns";
import { EmployeesTable } from "@/tables/employees/table";
import { columns as studentColumns } from "@/tables/students/columns";
import { StudentsTable } from "@/tables/students/table";
import { columns as teacherColumns } from "@/tables/teachers/columns";
import { TeachersTable } from "@/tables/teachers/table";
import { FC } from "react";

const UsersPage: FC = () => {
  const { teachers, students, employees } = useApi();

  return (
    <section>
      <h1>Сотрудники</h1>
      <EmployeesTable columns={columns} data={employees} />
      <h1>Преподаватели</h1>
      <TeachersTable columns={teacherColumns} data={teachers} />
      <h1>Студенты</h1>
      <StudentsTable columns={studentColumns} data={students} />
    </section>
  );
};

export default UsersPage;
