import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainTeacherCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.user.create({
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
        chats: {
          create: {
            name: "Новости",
            type: "USER_NEWS",
          },
        },
        profile: {
          create: {
            fio:
              data["surname"] +
              " " +
              data["firstname"] +
              " " +
              data["middlename"],
            role: data["role"],
            description: "Пользователь социальной сети YStudent",
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.teacher.created`,
      Buffer.from(JSON.stringify(response)),
    );
    console.log("teacher was created");
    ch.ack(msg);
  } catch {
    console.log("smth went wrong");
  }
};
