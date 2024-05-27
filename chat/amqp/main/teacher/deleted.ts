import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainTeacherDeleted = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const entity = await prisma.teacher.delete({
      where: {
        id: data["id"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.teacher.deleted`,
      Buffer.from(JSON.stringify(entity)),
    );
    ch.ack(msg);
    console.log("teacher was deleted");
  } catch {
    console.log("smth went wrong");
  }
};
