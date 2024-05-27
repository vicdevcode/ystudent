import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainTeacherUpdated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const entity = await prisma.user.update({
      where: {
        id: data["user"]["id"],
      },
      data: {
        firstname: data["firstname"],
        middlename: data["middlename"],
        surname: data["surname"],
        email: data["email"],
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.teacher.updated`,
      Buffer.from(JSON.stringify(entity)),
    );
    console.log("teacher was updated");
    ch.ack(msg);
  } catch {
    console.log("smth went wrong");
  }
};
