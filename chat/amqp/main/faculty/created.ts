import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainFacultyCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.faculty.create({
      data: {
        id: data["id"],
        name: data["name"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.faculty.created`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("faculty was created");
  } catch {
    console.log("smth went wrong");
  }
};
