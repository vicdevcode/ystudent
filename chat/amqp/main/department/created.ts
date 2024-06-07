import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainDepartmentCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const group = await prisma.department.create({
      data: {
        id: data["id"],
        name: data["name"],
        chat: {
          create: {
            name: data["name"],
            type: "NEWS",
          },
        },
        faculty: {
          connect: {
            id: data["faculty_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.department.created`,
      Buffer.from(JSON.stringify(group)),
    );
    ch.ack(msg);
    console.log("department was created");
  } catch (e) {
    console.log("smth went wrong", e);
  }
};
