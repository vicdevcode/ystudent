import "dotenv/config";
import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import amqplib from "amqplib/callback_api";
import { start } from "./amqp";
import cors from "cors";
import { amqpConfig, http_port } from "./config";
import { JwtUser } from "./jwt";
import { JwtPayload } from "jsonwebtoken";
import {
  addAdmin,
  chatAddMember,
  createChat,
  deleteAdmin,
  getAdmins,
  getAllChats,
  getCandidates,
} from "./controller/chat";
import {
  changeProfile,
  createTag,
  deleteTag,
  getProfile,
  search,
} from "./controller/profile";
import { getUserPayload, checkAuth } from "./middleware";
import { addImportant } from "./controller/important";
import { socketImp } from "./socket";
import { getAllUsers } from "./controller/user";

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

io.use(checkAuth);
app.use(express.json());
app.use(cors());
app.use(getUserPayload);

app.get("/api/v1/chat/candidates", getCandidates);

app.get("/api/v1/chat/get-all-chats", getAllChats);

app.get("/api/v1/chat/get-all", getAllUsers);

app.get("/api/v1/chat/get-profile", getProfile);

app.post("/api/v1/chat/add", chatAddMember);

app.post("/api/v1/chat/", createChat);

app.post("/api/v1/chat/search", search);

app.post("/api/v1/chat/delete-tag", deleteTag);

app.post("/api/v1/chat/change-profile", changeProfile);

app.post("/api/v1/chat/important", addImportant);

app.post("/api/v1/chat/create-tag", createTag);

app.post("/api/v1/chat/add-admin", addAdmin);

app.post("/api/v1/chat/delete-admin", deleteAdmin);

app.post("/api/v1/chat/get-admins", getAdmins);

io.on("connection", socketImp);

amqplib.connect(amqpConfig.url, start);

server.listen(http_port, () => {
  console.log("server running at " + process.env.HTTP_PORT);
});
