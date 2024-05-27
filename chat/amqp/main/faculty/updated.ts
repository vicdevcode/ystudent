import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainFacultyUpdated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.faculty.update({
      where: {
        id: data["id"],
      },
      data: {
        name: data["name"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.faculty.updated`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("faculty was updated");
  } catch {
    console.log("smth went wrong");
  }
};
