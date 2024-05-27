import { useApi } from "@/providers/api";
import { columns as studentColumns } from "@/tables/students/columns";
import { StudentsTable } from "@/tables/students/table";
import { columns as teacherColumns } from "@/tables/teachers/columns";
import { TeachersTable } from "@/tables/teachers/table";
import { FC } from "react";

const UsersPage: FC = () => {
  const { teachers, students } = useApi();

  return (
    <section>
      <h1>Преподаватели</h1>
      <TeachersTable columns={teacherColumns} data={teachers} />
      <h1>Студенты</h1>
      <StudentsTable columns={studentColumns} data={students} />
    </section>
  );
};

export default UsersPage;
