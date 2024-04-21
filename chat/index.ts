import "dotenv/config";
import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import amqplib from "amqplib/callback_api";
import { start } from "./amqp";
import cors from "cors";
import { amqpConfig, auth_check, http_port } from "./config";

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
  if (!access_token) next(err);
  const isAuth = await fetch(auth_check, {
    method: "GET",
    headers: {
      Authorization: "Bearer " + access_token,
    },
  }).then((res) => res.status === 200);
  if (isAuth) next();
  next(err);
});

app.use(cors());

app.get("/", (req, res) => {
  res.json({ hello: "world" });
});

io.on("connection", (socket) => {
  console.log("a user connected", socket.id);
});

amqplib.connect(amqpConfig.url, start);

server.listen(http_port, () => {
  console.log("server running at http://localhost:" + process.env.HTTP_PORT);
});
