import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainDepartmentCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const departmentData = JSON.parse(msg.content.toString());
    console.log(departmentData);
    const group = await prisma.department.create({
      data: {
        id: departmentData["id"],
        name: departmentData["name"],
        faculty: {
          connect: {
            id: departmentData["faculty_id"],
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
  } catch {
    console.log("smth went wrong");
  }
};
