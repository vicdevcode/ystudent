import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const AuthTeacherCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const teacherData = JSON.parse(msg.content.toString());
    const teacher = await prisma.user.create({
      data: {
        id: teacherData["id"],
        firstname: teacherData["firstname"],
        middlename: teacherData["middlename"],
        surname: teacherData["surname"],
        email: teacherData["email"],
        roleType: teacherData["role"],
        teacher: {
          create: {
            id: teacherData["role_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.teacher.created`,
      Buffer.from(JSON.stringify(teacher)),
    );
    console.log("teacher was created");
    ch.ack(msg);
  } catch {
    console.log("smth went wrong");
  }
};
