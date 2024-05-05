import amqplib from "amqplib/callback_api";
import { rabbitmq as r } from ".";
import AuthRouter from "./auth";

const router = async (msg: amqplib.Message | null) => {
  if (r.ch == undefined || msg == null) return;
  AuthRouter(r.ch, msg);
};

export default router;
