import { Request, Response } from "express";
import { InstanceOfJwt } from "../jwt";
import { prisma } from "../prisma";

export const getCandidates = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.user.findMany({
    include: {
      profile: {
        select: {
          id: true,
          fio: true,
          role: true,
          description: true,
          tags: true,
        },
      },
      chats: {
        select: {
          id: true,
          name: true,
          members: true,
          messages: true,
        },
      },
    },
  });
  res.json(user);
};

export const getAllChats = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.chat.findMany({
    where: {
      type: "OFFICIAL",
    },
  });
  res.json(user);
};

export const chatAddMember = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);
  if (typeof body["id"] !== "string") return res.status(400).send();
  if (typeof body["user_id"] !== "string") return res.status(400).send();

  const user = await prisma.user.findUnique({
    where: {
      id: body.user_id,
    },
  });
  if (!user) return res.sendStatus(400);

  const chat = await prisma.chat.update({
    where: {
      id: body.id,
    },
    include: {
      members: true,
      messages: true,
    },
    data: {
      members: {
        connect: {
          id: user.id,
        },
      },
    },
  });

  return res.status(200).json(chat);
};

export const createChat = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.user.findUnique({
    where: {
      email: req.user["email"],
    },
  });
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["name"] !== "string") return res.status(400).send();
  if (typeof body["type"] !== "string") return res.status(400).send();

  const chat = await prisma.chat.create({
    data: {
      name: body.name,
      type: body.type,
      members: {
        connect: {
          email: user?.email,
        },
      },
    },
  });

  console.log(chat);

  res.json(chat);
};

export const checkAdmin = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["chat_id"] !== "string") return res.status(400).send();
  if (typeof body["user_id"] !== "string") return res.status(400).send();

  const chat = await prisma.chatOnAdmins.findUnique({
    where: {
      chatId_userId: {
        chatId: body["chat_id"],
        userId: body["user_id"],
      },
    },
  });
  res.status(200).json(chat);
};
export const addAdmin = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["id"] !== "string") return res.status(400).send();
  if (typeof body["user_id"] !== "string") return res.status(400).send();
  if (typeof body["type"] !== "string") return res.status(400).send();

  let chatId = "";
  let userId = "";
  const checkUser = await prisma.user.findUnique({
    where: {
      id: body["user_id"],
    },
  });
  if (body["type"] === "faculty") {
    const faculty = await prisma.faculty.findUnique({
      where: {
        id: body["id"],
      },
    });
    if (faculty) chatId = faculty.chatId as string;
    if (checkUser?.roleType == "EMPLOYEE") userId = checkUser.id;
  }
  if (body["type"] === "department") {
    const department = await prisma.department.findUnique({
      where: {
        id: body["id"],
      },
    });
    if (department) chatId = department.chatId as string;
    if (checkUser?.roleType == "EMPLOYEE" || checkUser?.roleType == "TEACHER")
      userId = checkUser.id;
  }
  if (chatId === "" || userId === "") return res.status(400).send();
  const chat = await prisma.chatOnAdmins.create({
    data: {
      chatId: chatId,
      userId: userId,
    },
  });
  res.json(chat);
};

export const deleteAdmin = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["id"] !== "string") return res.status(400).send();
  if (typeof body["user_id"] !== "string") return res.status(400).send();
  if (typeof body["type"] !== "string") return res.status(400).send();

  let chatId = "";
  let userId = "";
  const checkUser = await prisma.user.findUnique({
    where: {
      id: body["user_id"],
    },
  });
  if (body["type"] === "faculty") {
    const faculty = await prisma.faculty.findUnique({
      where: {
        id: body["id"],
      },
    });
    if (faculty) chatId = faculty.chatId as string;
    if (checkUser?.roleType == "EMPLOYEE") userId = checkUser.id;
  }
  if (body["type"] === "department") {
    const department = await prisma.department.findUnique({
      where: {
        id: body["id"],
      },
    });
    if (department) chatId = department.chatId as string;
    if (checkUser?.roleType == "EMPLOYEE" || checkUser?.roleType == "TEACHER")
      userId = checkUser.id;
  }
  if (chatId === "" || userId === "") return res.status(400).send();
  const chat = await prisma.chatOnAdmins.delete({
    where: {
      chatId_userId: {
        userId: userId,
        chatId: chatId,
      },
    },
  });
  res.json(chat);
};

export const getAdmins = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["id"] !== "string") return res.status(400).send();
  if (typeof body["type"] !== "string") return res.status(400).send();

  let chatId = "";
  if (body["type"] === "faculty") {
    const faculty = await prisma.faculty.findUnique({
      where: {
        id: body["id"],
      },
    });

    if (!faculty?.chatId) return res.status(400).send();
    chatId = faculty.chatId;
  } else if (body["type"] === "department") {
    const department = await prisma.department.findUnique({
      where: {
        id: body["id"],
      },
    });

    if (!department?.chatId) return res.status(400).send();
    chatId = department.chatId;
  }

  const admins = await prisma.chatOnAdmins.findMany({
    where: {
      chatId: chatId,
    },
  });

  const users = await prisma.user.findMany({
    where: {
      id: { in: admins.map((a) => a.userId) },
    },
  });
  res.json(users);
};
