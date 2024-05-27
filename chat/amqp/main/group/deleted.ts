import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainGroupDeleted = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.group.delete({
      where: {
        id: data["id"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.group.deleted`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("group was deleted");
  } catch {
    console.log("smth went wrong");
  }
};
