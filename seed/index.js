const fio = require("./fio.json");

api = process.env.API;

const randomNumber = (min, max) => {
  return Math.floor(Math.random() * (max - min) + min);
};

let accessToken = "";

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
  const res = await fetch(api + "/faculty", {
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

const addGroup = async (faculty_id, group_name) => {
  const res = await fetch(api + "/group", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    },
    body: JSON.stringify({
      faculty_id: faculty_id,
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

  const res = await fetch(api + "/student/create-with-user", {
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

  const res = await fetch(api + "/teacher/create-with-user", {
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

const start = async () => {
  const faculties = await checkFaculties();

  if (faculties.length == 0) {
    await authorize();
    const imi = await addFaculty();
    const t1 = await addTeacher(`t1@gmail.com`);
    const t2 = await addTeacher(`t2@gmail.com`);
    for (let i = 3; i <= 5; i++) {
      await addTeacher(`t${i}@gmail.com`);
    }
    const group_ivt202 = await addGroup(imi["id"], "ИВТ-20-2");
    await updateCuratorGroup(group_ivt202["id"], t1["teacher"]["id"]);
    for (let i = 1; i <= 20; i++) {
      await addStudent(`s${i}@gmail.com`, group_ivt202["id"]);
    }
    const group_ivt201 = await addGroup(imi["id"], "ИВТ-20-1");
    await updateCuratorGroup(group_ivt201["id"], t2["teacher"]["id"]);
    for (let i = 21; i <= 40; i++) {
      await addStudent(`s${i}@gmail.com`, group_ivt201["id"]);
    }
  }
};

start();
