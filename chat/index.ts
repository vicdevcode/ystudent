import "dotenv/config";
import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import amqplib from "amqplib/callback_api";
import { start } from "./amqp";
import { amqp_url } from "./amqp/constant";

const app = express();
const server = createServer(app);
const io = new Server(server);

app.get("/", (req, res) => {
  res.json({ hello: "world" });
});

io.on("connection", (socket) => {
  console.log("a user connected");
});

amqplib.connect(amqp_url, start);

server.listen(process.env.HTTP_PORT, () => {
  console.log("server running at http://localhost:" + process.env.HTTP_PORT);
});
