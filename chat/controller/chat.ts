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
