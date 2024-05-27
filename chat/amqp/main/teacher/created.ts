import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainTeacherCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const entity = await prisma.user.create({
      data: {
        id: data["id"],
        firstname: data["firstname"],
        middlename: data["middlename"],
        surname: data["surname"],
        email: data["email"],
        roleType: data["role"],
        teacher: {
          create: {
            id: data["role_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.teacher.created`,
      Buffer.from(JSON.stringify(entity)),
    );
    console.log("teacher was created");
    ch.ack(msg);
  } catch {
    console.log("smth went wrong");
  }
};
