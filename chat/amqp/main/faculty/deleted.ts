import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainFacultyDeleted = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.faculty.delete({
      where: {
        id: data["id"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.faculty.deleted`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("faculty was deleted");
  } catch {
    console.log("smth went wrong");
  }
};
