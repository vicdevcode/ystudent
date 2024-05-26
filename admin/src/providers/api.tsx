import { useToast } from "@/components/ui/use-toast";
import {
  CreateDepartment,
  CreateFaculty,
  CreateGroup,
  CreateTeacher,
  EditDepartment,
  EditFaculty,
  EditGroup,
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

interface ApiContext {
  faculties: Faculties[];
  departments: Departments[];
  groups: Groups[];
  teachers: Teachers[];
  getFaculties: (token: string) => Promise<void>;
  getDepartments: (token: string) => Promise<void>;
  getGroups: (token: string) => Promise<void>;
  getTeachers: (token: string) => Promise<void>;
  createTeacher: (token: string, data: CreateTeacher) => Promise<void>;
  createGroup: (token: string, data: CreateGroup) => Promise<void>;
  createDepartment: (token: string, data: CreateDepartment) => Promise<void>;
  createFaculty: (token: string, data: CreateFaculty) => Promise<void>;
  editTeacher: (token: string, data: EditTeacher) => Promise<void>;
  editGroup: (token: string, data: EditGroup) => Promise<void>;
  editDepartment: (token: string, data: EditDepartment) => Promise<void>;
  editFaculty: (token: string, data: EditFaculty) => Promise<void>;
  deleteTeacher: (token: string, id: string) => Promise<void>;
  deleteGroup: (token: string, id: string) => Promise<void>;
  deleteDepartment: (token: string, id: string) => Promise<void>;
  deleteFaculty: (token: string, id: string) => Promise<void>;
}

const defaultValues: ApiContext = {
  faculties: [],
  departments: [],
  groups: [],
  teachers: [],
  getFaculties: async () => { },
  getDepartments: async () => { },
  getGroups: async () => { },
  getTeachers: async () => { },
  createTeacher: async () => { },
  createGroup: async () => { },
  createDepartment: async () => { },
  createFaculty: async () => { },
  editTeacher: async () => { },
  editGroup: async () => { },
  editDepartment: async () => { },
  editFaculty: async () => { },
  deleteTeacher: async () => { },
  deleteGroup: async () => { },
  deleteDepartment: async () => { },
  deleteFaculty: async () => { },
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
      case "/teachers":
        getTeachers(token);
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
          email: json["teachers"][i]["user"]["email"],
        });
      }
      setTeachers(data);
    });
  };

  const createTeacher = async (token: string) => {
    if (!token) return;
    console.log("impliment me");
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

  const editTeacher = async (token: string) => {
    if (!token) return;
    console.log("impliment me");
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

  const exposed = {
    faculties,
    departments,
    groups,
    teachers,
    getFaculties,
    getDepartments,
    getGroups,
    getTeachers,
    createTeacher,
    createGroup,
    createDepartment,
    createFaculty,
    editTeacher,
    editGroup,
    editDepartment,
    editFaculty,
    deleteTeacher,
    deleteGroup,
    deleteDepartment,
    deleteFaculty,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export default ApiProvider;

export const useApi = () => useContext(Context);
