import { useToast } from "@/components/ui/use-toast";
import {
  CreateDepartment,
  CreateFaculty,
  CreateGroup,
  CreateStudent,
  CreateTeacher,
  EditDepartment,
  EditFaculty,
  EditGroup,
  EditStudent,
  EditTeacher,
} from "@/lib/types";
import { Departments } from "@/tables/departments/columns";
import { Faculties } from "@/tables/faculties/columns";
import { Groups } from "@/tables/groups/columns";
import { Teachers } from "@/tables/teachers/columns";
import {
  FC,
  PropsWithChildren,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import { useAuth } from "./auth";
import { useLocation } from "react-router-dom";
import { Students } from "@/tables/students/columns";

interface ApiContext {
  faculties: Faculties[];
  departments: Departments[];
  groups: Groups[];
  teachers: Teachers[];
  students: Students[];
  getFaculties: (token: string) => Promise<void>;
  getDepartments: (token: string) => Promise<void>;
  getGroups: (token: string) => Promise<void>;
  getTeachers: (token: string) => Promise<void>;
  getStudents: (token: string) => Promise<void>;
  createFaculty: (token: string, data: CreateFaculty) => Promise<void>;
  createDepartment: (token: string, data: CreateDepartment) => Promise<void>;
  createGroup: (token: string, data: CreateGroup) => Promise<void>;
  createTeacher: (token: string, data: CreateTeacher) => Promise<void>;
  createStudent: (token: string, data: CreateStudent) => Promise<void>;
  editFaculty: (token: string, data: EditFaculty) => Promise<void>;
  editDepartment: (token: string, data: EditDepartment) => Promise<void>;
  editGroup: (token: string, data: EditGroup) => Promise<void>;
  editTeacher: (token: string, data: EditTeacher) => Promise<void>;
  editStudent: (token: string, data: EditStudent) => Promise<void>;
  deleteFaculty: (token: string, id: string) => Promise<void>;
  deleteDepartment: (token: string, id: string) => Promise<void>;
  deleteGroup: (token: string, id: string) => Promise<void>;
  deleteTeacher: (token: string, id: string) => Promise<void>;
  deleteStudent: (token: string, id: string) => Promise<void>;
}

const defaultValues: ApiContext = {
  faculties: [],
  departments: [],
  groups: [],
  teachers: [],
  students: [],
  getFaculties: async () => { },
  getDepartments: async () => { },
  getGroups: async () => { },
  getTeachers: async () => { },
  getStudents: async () => { },
  createFaculty: async () => { },
  createDepartment: async () => { },
  createGroup: async () => { },
  createTeacher: async () => { },
  createStudent: async () => { },
  editFaculty: async () => { },
  editDepartment: async () => { },
  editGroup: async () => { },
  editTeacher: async () => { },
  editStudent: async () => { },
  deleteFaculty: async () => { },
  deleteDepartment: async () => { },
  deleteGroup: async () => { },
  deleteTeacher: async () => { },
  deleteStudent: async () => { },
};

const Context = createContext(defaultValues);

interface ApiProvidersProps extends PropsWithChildren { }

const ApiProvider: FC<ApiProvidersProps> = ({ children }) => {
  const { token } = useAuth();
  const location = useLocation();
  const [faculties, setFaculties] = useState<Faculties[]>(
    defaultValues.faculties,
  );
  const [departments, setDepartments] = useState<Departments[]>(
    defaultValues.departments,
  );
  const [groups, setGroups] = useState<Groups[]>(defaultValues.groups);
  const [teachers, setTeachers] = useState<Teachers[]>(defaultValues.teachers);
  const [students, setStudents] = useState<Students[]>(defaultValues.students);
  const { toast } = useToast();

  useEffect(() => {
    switch (location.pathname) {
      case "/faculties":
        getFaculties(token);
        break;
      case "/departments":
        getDepartments(token);
        getFaculties(token);
        break;
      case "/groups":
        getGroups(token);
        getDepartments(token);
        getTeachers(token);
        break;
      case "/users":
        getGroups(token);
        getTeachers(token);
        getStudents(token);
        break;
    }
  }, [token, location]);

  const getFaculties = async (token: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/faculties", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) return;
      const data = [];
      for (let i = 0; i < json["faculties"].length; i++) {
        data.push({
          id: json["faculties"][i]["id"],
          name: json["faculties"][i]["name"],
        });
      }
      setFaculties(data);
    });
  };

  const getDepartments = async (token: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/departments", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) return;
      const data = [];
      for (let i = 0; i < json["departments"].length; i++) {
        data.push({
          id: json["departments"][i]["id"],
          name: json["departments"][i]["name"],
          faculty_name: json["departments"][i]["faculty_name"],
          faculty_id: json["departments"][i]["faculty_id"],
        });
      }
      setDepartments(data);
    });
  };

  const getGroups = async (token: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/groups", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) return;
      const data = [];
      for (let i = 0; i < json["groups"].length; i++) {
        data.push({
          id: json["groups"][i]["id"],
          name: json["groups"][i]["name"],
          department_name: json["groups"][i]["department_name"],
          department_id: json["groups"][i]["department_id"],
          curator_id: json["groups"][i]["curator_id"],
          curator_fio: json["groups"][i]["curator_fio"],
        });
      }
      setGroups(data);
    });
  };

  const getTeachers = async (token: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/teachers", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) return;
      const data = [];
      for (let i = 0; i < json["teachers"].length; i++) {
        data.push({
          id: json["teachers"][i]["id"],
          name: `${json["teachers"][i]["user"]["surname"]} ${json["teachers"][i]["user"]["firstname"]} ${json["teachers"][i]["user"]["middlename"]}`,
          firstname: json["teachers"][i]["user"]["firstname"],
          surname: json["teachers"][i]["user"]["surname"],
          middlename: json["teachers"][i]["user"]["middlename"],
          email: json["teachers"][i]["user"]["email"],
        });
      }
      setTeachers(data);
    });
  };

  const getStudents = async (token: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/students", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) return;
      const data = [];
      for (let i = 0; i < json["students"].length; i++) {
        data.push({
          id: json["students"][i]["id"],
          name: `${json["students"][i]["user"]["surname"]} ${json["students"][i]["user"]["firstname"]} ${json["students"][i]["user"]["middlename"]}`,
          firstname: json["students"][i]["user"]["firstname"],
          surname: json["students"][i]["user"]["surname"],
          middlename: json["students"][i]["user"]["middlename"],
          email: json["students"][i]["user"]["email"],
          group_name: json["students"][i]["group_name"],
          group_id: json["students"][i]["group_id"],
        });
      }
      setStudents(data);
    });
  };

  const createFaculty = async (token: string, data: CreateFaculty) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/faculty/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(data),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getFaculties(token);
      toast({
        title: "Факультет был создан",
      });
    });
  };

  const createDepartment = async (token: string, data: CreateDepartment) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/department/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(data),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getDepartments(token);
      toast({
        title: "Кафедра была создана",
      });
    });
  };

  const createGroup = async (token: string, data: CreateGroup) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/group/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(data),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getGroups(token);
      toast({
        title: "Группа была создана",
      });
    });
  };

  const createTeacher = async (token: string, data: CreateTeacher) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/teacher/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(data),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getTeachers(token);
      toast({
        title: "Преподаватель был создан",
        description: "Пароль: " + json["password"],
      });
    });
  };

  const createStudent = async (token: string, data: CreateStudent) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/student/", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify(data),
    }).then(async (res) => {
      const json = await res.json();
      if (res.status != 200) {
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getStudents(token);
      toast({
        title: "Студент был создан",
        description: "Пароль: " + json["password"],
      });
    });
  };

  const editFaculty = async (token: string, data: EditFaculty) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/faculty/" + data.id, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        name: data.name,
      }),
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getFaculties(token);
      toast({
        title: "Факультет успешно отредактирован",
      });
    });
  };

  const editDepartment = async (token: string, data: EditDepartment) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/department/" + data.id, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        name: data.name,
        faculty_id: data.faculty_id,
      }),
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getDepartments(token);
      toast({
        title: "Кафедра успешно отредактирована",
      });
    });
  };

  const editGroup = async (token: string, data: EditGroup) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/group/" + data.id, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        name: data.name,
        department_id: data.department_id,
        curator_id: data.curator_id,
      }),
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getGroups(token);
      toast({
        title: "Группа успешно отредактирована",
      });
    });
  };

  const editTeacher = async (token: string, data: EditTeacher) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/teacher/" + data.id, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        firstname: data.firstname,
        middlename: data.middlename,
        surname: data.surname,
        email: data.email,
      }),
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getTeachers(token);
      toast({
        title: "Преподаватель успешно отредактирован",
      });
    });
  };

  const editStudent = async (token: string, data: EditStudent) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/student/" + data.id, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        firstname: data.firstname,
        middlename: data.middlename,
        surname: data.surname,
        email: data.email,
        group_id: data.group_id,
      }),
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getStudents(token);
      toast({
        title: "Студент успешно отредактирован",
      });
    });
  };

  const deleteFaculty = async (token: string, id: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/faculty/" + id, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        console.log(json);
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getFaculties(token);
      toast({
        title: "Факультет успешно удален",
      });
    });
  };

  const deleteDepartment = async (token: string, id: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/department/" + id, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getDepartments(token);
      toast({
        title: "Кафедра успешна удалена",
      });
    });
  };

  const deleteGroup = async (token: string, id: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/group/" + id, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getGroups(token);
      toast({
        title: "Группа успешно удалена",
      });
    });
  };

  const deleteTeacher = async (token: string, id: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/teacher/" + id, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getTeachers(token);
      toast({
        title: "Преподаватель успешно удален",
      });
    });
  };

  const deleteStudent = async (token: string, id: string) => {
    if (!token) return;
    return fetch(import.meta.env.VITE_MAIN_API + "/student/" + id, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + token,
      },
    }).then(async (res) => {
      if (res.status != 200) {
        const json = await res.json();
        toast({
          variant: "destructive",
          title: "ERROR",
          description: json["message"],
        });
        return;
      }
      getStudents(token);
      toast({
        title: "Студент успешно удален",
      });
    });
  };

  const exposed = {
    faculties,
    departments,
    groups,
    teachers,
    students,
    getFaculties,
    getDepartments,
    getGroups,
    getTeachers,
    getStudents,
    createFaculty,
    createDepartment,
    createGroup,
    createTeacher,
    createStudent,
    editFaculty,
    editDepartment,
    editGroup,
    editTeacher,
    editStudent,
    deleteFaculty,
    deleteDepartment,
    deleteGroup,
    deleteTeacher,
    deleteStudent,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export default ApiProvider;

export const useApi = () => useContext(Context);
