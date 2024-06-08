export interface EditFaculty {
  id: string;
  name: string;
}

export interface EditDepartment {
  id: string;
  name?: string;
  faculty_id?: string;
}

export interface EditGroup {
  id: string;
  name?: string;
  department_id?: string;
  curator_id?: string;
}

export interface EditTeacher {
  id: string;
  firstname?: string;
  middlename?: string;
  surname?: string;
  email?: string;
}

export interface EditEmployee {
  id: string;
  firstname?: string;
  middlename?: string;
  surname?: string;
  email?: string;
}

export interface EditStudent {
  id: string;
  firstname?: string;
  middlename?: string;
  surname?: string;
  email?: string;
  group_id: string;
}

export interface CreateFaculty {
  name: string;
}

export interface CreateDepartment {
  name: string;
  faculty_id?: string;
}

export interface CreateGroup {
  name: string;
  department_id: string;
  curator_id: string;
}

export interface CreateTeacher {
  firstname: string;
  middlename?: string;
  surname: string;
  email: string;
}

export interface CreateEmployee {
  firstname: string;
  middlename?: string;
  surname: string;
  email: string;
}

export interface CreateStudent {
  firstname: string;
  middlename?: string;
  surname: string;
  email: string;
  group_id: string;
}

export interface CreateChatAdmin {
  id: string;
  user_id: string;
  type: "faculty" | "department";
}

export interface GetChatAdmins {
  id: string;
  type: "faculty" | "department";
}
