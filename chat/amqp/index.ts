import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../config";
import router from "./router";

export interface RabbitMQ {
  conn?: amqplib.Connection;
  ch?: amqplib.Channel;
}

export const rabbitmq: RabbitMQ = {};

export function start(err: Error, conn: amqplib.Connection) {
  if (err) throw err;

  rabbitmq.conn = conn;

  conn.createChannel((err: Error, ch: amqplib.Channel) => {
    if (err) throw err;

    rabbitmq.ch = ch;

    ch.assertExchange(amqpConfig.exchange, "topic", {
      durable: false,
    });

    ch.assertQueue(
      amqpConfig.queue_name,
      {
        durable: true,
      },
      (err: Error, q: amqplib.Replies.AssertQueue) => {
        if (err) throw err;

        amqpConfig.topics.map((topic) =>
          ch.bindQueue(q.queue, amqpConfig.exchange, topic),
        );

        ch.consume(q.queue, router, {
          noAck: false,
        });
      },
    );
    console.log("rabbitmq started");
  });
}
