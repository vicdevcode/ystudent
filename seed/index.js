const fio = require("./fio.json");

api = process.env.API;

const randomNumber = (min, max) => {
  return Math.floor(Math.random() * (max - min) + min);
};

let accessToken =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlmNWIzYWRkLWQ2NDUtNDJiNS1iMzkxLWY0NWE0ZmFhN2QyOCIsImVtYWlsIjoidmljZGV2Y29kZUBnbWFpbC5jb20iLCJyb2xlIjoiQURNSU4iLCJleHAiOjE3MTkzNjUxNTh9.JYqRjtzK7mW70CYY68qKk2SYo_F-kFlIGwa0x8WHl_c";

const authorize = async () => {
  const res = await fetch(api + "/auth/admin", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      login: process.env.LOGIN,
      password: process.env.PASSWORD,
    }),
  });

  const json = await res.json();

  accessToken = json["access_token"];

  return json;
};

const generateFio = () => {
  if (randomNumber(0, 2) != 1) {
    return {
      firstname:
        fio["male_firstnames"][randomNumber(0, fio["male_firstnames"].length)],
      surname:
        fio["male_surnames"][randomNumber(0, fio["male_surnames"].length)],
      middlename:
        fio["male_middlenames"][
        randomNumber(0, fio["male_middlenames"].length)
        ],
    };
  }
  return {
    firstname:
      fio["female_firstnames"][
      randomNumber(0, fio["female_firstnames"].length)
      ],
    surname:
      fio["female_surnames"][randomNumber(0, fio["female_surnames"].length)],
    middlename:
      fio["female_middlenames"][
      randomNumber(0, fio["female_middlenames"].length)
      ],
  };
};

const checkFaculties = async () => {
  const res = await fetch(api + "/faculties", {
    method: "GET",
  });
  const json = await res.json();

  return json["faculties"];
};

const addFaculty = async () => {
  const res = await fetch(api + "/faculty", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      name: "ИМИ",
    }),
  });

  const json = await res.json();

  return json;
};

const addGroup = async (faculty_id, group_name, curator_id) => {
  const res = await fetch(api + "/group", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      department_id: faculty_id,
      curator_id: curator_id,
      name: group_name,
    }),
  });

  const json = await res.json();

  return json;
};

const updateCuratorGroup = async (group_id, teacher_id) => {
  const res = await fetch(api + "/group/" + group_id, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      teacher_id: teacher_id,
    }),
  });

  const json = await res.json();

  return json;
};

const addStudent = async (email, group_id) => {
  const fio = generateFio();

  const res = await fetch(api + "/student", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      firstname: fio.firstname,
      surname: fio.surname,
      middlename: fio.middlename,
      email: email,
      group_id: group_id,
    }),
  });

  const json = await res.json();

  return json;
};

const addTeacher = async (email) => {
  const fio = generateFio();

  const res = await fetch(api + "/teacher", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      firstname: fio.firstname,
      surname: fio.surname,
      middlename: fio.middlename,
      email: email,
    }),
  });

  const json = await res.json();

  return json;
};

const addDepartment = async (faculty_id, name) => {
  const res = await fetch(api + "/department", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      faculty_id: faculty_id,
      name: name,
    }),
  });

  const json = await res.json();

  return json;
};

const start = async () => {
  const faculties = await checkFaculties();
  if (faculties.length == 0) {
    const imi = await addFaculty();
    const it = await addDepartment(imi["id"], "ИТ");
    const t1 = await addTeacher(`t1@gmail.com`);
    const t2 = await addTeacher(`t2@gmail.com`);
    for (let i = 3; i <= 5; i++) {
      await addTeacher(`t${i}@gmail.com`);
    }
    const group_ivt202 = await addGroup(it["id"], "ИВТ-20-2", t1["id"]);
    for (let i = 1; i <= 20; i++) {
      await addStudent(`s${i}@gmail.com`, group_ivt202["id"]);
    }
    const group_ivt201 = await addGroup(it["id"], "ИВТ-20-1", t2["id"]);
    for (let i = 21; i <= 40; i++) {
      await addStudent(`s${i}@gmail.com`, group_ivt201["id"]);
    }
  }
};

start();
