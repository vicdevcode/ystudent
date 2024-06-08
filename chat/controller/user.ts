import { Request, Response } from "express";
import { InstanceOfJwt } from "../jwt";
import { prisma } from "../prisma";

export const getAllUsers = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.user.findUnique({
    where: {
      email: req.user["email"],
    },
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
          type: true,
          members: true,
          messages: true,
        },
      },
    },
  });
  res.status(200).json(user);
};
