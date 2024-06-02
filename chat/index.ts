import "dotenv/config";
import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import amqplib from "amqplib/callback_api";
import { start } from "./amqp";
import cors from "cors";
import { amqpConfig, auth_check, http_port } from "./config";
import { prisma } from "./prisma";
import { InstanceOfJwt, JwtUser, getPayload } from "./jwt";
import { JwtPayload } from "jsonwebtoken";
import { ChatType } from "@prisma/client";

declare global {
  namespace Express {
    export interface Request {
      user: JwtUser | JwtPayload | string;
    }
  }
}

const app = express();
const server = createServer(app);
const io = new Server(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"],
    credentials: true,
  },
});

io.use(async (socket, next) => {
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
});

app.use(express.json());
app.use(cors());

app.use((req, res, next) => {
  const token = req.headers.authorization;
  if (!token) return res.status(401).send();

  const payload = getPayload(token.split(" ")[1]);

  if (!payload) return res.status(401).send();

  req.user = payload;

  next();
});

app.get("/", (req, res) => {
  res.json({ hello: "world" });
});

app.post("/api/v1/chat/", async (req, res) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.user.findUnique({
    where: {
      email: req.user["email"],
    },
  });
  const body = req.body;
  console.log(body);
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
});

app.get("/api/v1/chat/get-all", async (req, res) => {
  if (!InstanceOfJwt(req.user)) return res.status(401).send();
  const user = await prisma.user.findUnique({
    where: {
      email: req.user["email"],
    },
    include: {
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
  res.json({
    chats: user?.chats,
  });
});

type UserPayload = JwtUser | JwtPayload | string;

io.on("connection", (socket) => {
  console.log("a user connected", socket.id);
  const access_token = socket.handshake.auth.access_token;
  const payload: UserPayload = getPayload(access_token);

  socket.on("join_room", (chat_id) => {
    socket.join(chat_id);
    console.log(socket.id, chat_id);
    io.to(chat_id).emit("user_joined", {
      userId: socket.id,
      username: "User",
      room: chat_id,
    });
    socket.on("send_message", async (message) => {
      let data = {
        email: "",
      };
      if (InstanceOfJwt(payload)) data.email = payload["email"];
      const user = await prisma.user.findUnique({
        where: {
          email: data.email,
        },
      });
      if (!user?.id) return;
      const newMessage = await prisma.message.create({
        data: {
          chatId: chat_id,
          content: message,
          senderId: user?.id,
          senderFio:
            user.surname + " " + user.firstname + " " + user.middlename,
        },
      });
      socket.to(chat_id).emit("receive_message", {
        senderId: newMessage.senderId,
        senderFio: newMessage.senderFio,
        content: message,
        createdAt: newMessage.createdAt,
      });
    });
    socket.on("leave_room", () => {
      socket.leave(chat_id);
    });
  });
  socket.on("disconnect", (reason) => {
    console.log(`[${socket.id}] socket disconnected - ${reason}`);
  });
});

amqplib.connect(amqpConfig.url, start);

server.listen(http_port, () => {
  console.log("server running at http://localhost:" + process.env.HTTP_PORT);
});
