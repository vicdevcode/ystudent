import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainStudentCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const group = await prisma.group.findUnique({
      where: {
        id: data["group_id"],
      },
    });
    if (!group?.chatId) throw new Error("can not get group");
    const response = await prisma.user.create({
      data: {
        id: data["id"],
        firstname: data["firstname"],
        middlename: data["middlename"],
        surname: data["surname"],
        email: data["email"],
        roleType: data["role"],
        chats: {
          connect: {
            id: group.chatId,
          },
        },
        student: {
          create: {
            group: {
              connect: {
                id: data["group_id"],
              },
            },
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.student.created`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg as amqplib.Message);
    console.log("student was created");
  } catch (e) {
    console.log("smth went wrong", e);
  }
};