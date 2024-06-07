import { Socket } from "socket.io";
import { ExtendedError } from "socket.io/dist/namespace";
import { auth_check } from "./config";
import { NextFunction, Request, Response } from "express";
import { getPayload } from "./jwt";

export const checkAuth = async (
  socket: Socket,
  next: (err?: ExtendedError) => void,
) => {
  const access_token = socket.handshake.auth.access_token;
  const err = new Error("unathenticated");
  if (!access_token) return next(err);
  const isAuth = await fetch(auth_check, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + access_token,
    },
  }).then((res) => res.status === 200);
  if (isAuth) return next();
  next(err);
};

export const getUserPayload = (
  req: Request,
  res: Response,
  next: NextFunction,
) => {
  const token = req.headers.authorization;
  if (!token) return res.status(401).send();

  const payload = getPayload(token.split(" ")[1]);

  if (!payload) return res.status(401).send();

  req.user = payload;

  next();
};
