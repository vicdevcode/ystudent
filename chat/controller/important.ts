import { Request, Response } from "express";
import { InstanceOfJwt } from "../jwt";
import { prisma } from "../prisma";

export const addImportant = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);
  if (typeof body.id !== "string") return res.sendStatus(400);
  if (typeof body.important !== "boolean") return res.sendStatus(400);
  if (typeof body.is_news !== "boolean") return res.sendStatus(400);
  if (typeof body.is_task !== "boolean") return res.sendStatus(400);
  const message = await prisma.message.update({
    where: { id: body.id },
    data: {
      important: body.important,
      is_news: body.is_news,
      is_task: body.is_task,
    },
  });
  res.status(200).json(message);
};
