import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const AuthStudentCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const studentData = JSON.parse(msg.content.toString());
    const user = await prisma.user.create({
      data: {
        id: studentData["id"],
        firstname: studentData["firstname"],
        middlename: studentData["middlename"],
        surname: studentData["surname"],
        email: studentData["email"],
        roleType: studentData["role_type"],
        student: {
          create: {
            group: {
              connect: {
                id: studentData["student"]["group_id"],
              },
            },
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.student.created`,
      Buffer.from(JSON.stringify(user)),
    );
    ch.ack(msg as amqplib.Message);
    console.log("student was created");
  } catch {
    console.log("smth went wrong");
  }
};
