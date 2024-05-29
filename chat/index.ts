import "dotenv/config";
import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import amqplib from "amqplib/callback_api";
import { start } from "./amqp";
import cors from "cors";
import { amqpConfig, auth_check, http_port } from "./config";
import { prisma } from "./prisma";

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

app.get("/", (req, res) => {
  res.json({ hello: "world" });
});

app.post("/api/v1/chat/", async (req, res) => {
  const body = req.body;
  if (!body) return res.sendStatus(400);

  if (typeof body["name"] !== "string") return;
  if (typeof body["type"] !== "string") return;

  const chat = await prisma.chat.create({
    data: {
      name: body.name,
      type: body.type,
    },
  });

  console.log(chat);

  res.json(chat);
});

app.get("/api/v1/chat/get-all", async (req, res) => {
  const chats = await prisma.chat.findMany();
  res.json({
    chats: chats,
  });
});

io.on("connection", (socket) => {
  console.log("a user connected", socket.id);
  socket.on("joinRoom", (roomName) => {
    socket.join(roomName);
    console.log(socket.id, roomName);
    io.to(roomName).emit("userJoined", {
      userId: socket.id,
      username: "User",
      room: roomName,
    });
  });
  socket.on("disconnect", (reason) => {
    console.log(`[${socket.id}] socket disconnected - ${reason}`);
  });
});

setInterval(() => {
  io.sockets.emit("time-msg", { time: new Date().toISOString() });
}, 1000);

amqplib.connect(amqpConfig.url, start);

server.listen(http_port, () => {
  console.log("server running at http://localhost:" + process.env.HTTP_PORT);
});
