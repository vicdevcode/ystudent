import amqplib from "amqplib/callback_api";
import { rabbitmq as r } from ".";
import MainRouter from "./main";

const router = async (msg: amqplib.Message | null) => {
  if (r.ch == undefined || msg == null) return;
  MainRouter(r.ch, msg);
};

export default router;
