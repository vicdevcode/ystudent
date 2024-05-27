import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainDepartmentDeleted = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.department.delete({
      where: {
        id: data["id"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.department.deleted`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("department was deleted");
  } catch {
    console.log("smth went wrong");
  }
};
