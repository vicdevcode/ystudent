import { Request, Response } from "express";
import { InstanceOfJwt } from "../jwt";
import { prisma } from "../prisma";

export const search = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body && !body.tag) return res.sendStatus(400);
  const profile = await prisma.profile.findMany({
    where: {
      NOT: {
        userId: req.user.id,
      },
      tags: {
        some: {
          name: body.tag,
        },
      },
    },
    include: {
      tags: true,
    },
  });
  res.status(200).json(profile);
};

export const getProfile = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const profile = await prisma.profile.findUnique({
    where: {
      userId: req.user.id,
    },
    include: {
      tags: true,
    },
  });
  res.status(200).json(profile);
};

export const deleteTag = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body && !body.tag) return res.sendStatus(400);
  const profile = await prisma.profile.update({
    where: {
      userId: req.user.id,
    },
    include: {
      tags: true,
    },
    data: {
      description: body.description,
      tags: {
        delete: {
          name: body.tag,
        },
      },
    },
  });
  return res.status(200).json(profile);
};

export const changeProfile = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body && !body.tags && !body.description) return res.sendStatus(400);
  if (!(body.tags instanceof Array)) return res.sendStatus(400);
  const profile = await prisma.profile.update({
    where: {
      userId: req.user.id,
    },
    include: {
      tags: true,
    },
    data: {
      description: body.description,
      tags: {
        connectOrCreate: body.tags.map((tag: string) => {
          return {
            where: { name: tag },
            create: { name: tag },
          };
        }),
      },
    },
  });
  return res.status(200).json(profile);
};

export const createTag = async (req: Request, res: Response) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const body = req.body;
  if (!body) return res.sendStatus(400);
  if (typeof body.name !== "string") return res.sendStatus(400);
  const tag = await prisma.tag.findFirst({
    where: {
      name: body.name,
    },
  });
  if (!tag) {
    const newTag = await prisma.tag.create({
      data: {
        name: body.name,
      },
    });
    return res.status(200).json(newTag);
  }
  return res.status(200).json(tag);
};
