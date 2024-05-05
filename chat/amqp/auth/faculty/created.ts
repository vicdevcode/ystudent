import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const AuthFacultyCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const facultyData = JSON.parse(msg.content.toString());
    const faculty = await prisma.faculty.create({
      data: {
        id: facultyData["id"],
        name: facultyData["name"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.faculty.created`,
      Buffer.from(JSON.stringify(faculty)),
    );
    ch.ack(msg);
    console.log("faculty was created");
  } catch {
    console.log("smth went wrong");
  }
};
