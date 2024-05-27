import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainGroupCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const data = JSON.parse(msg.content.toString());
    const response = await prisma.group.create({
      data: {
        id: data["id"],
        name: data["name"],
        department: {
          connect: {
            id: data["department_id"],
          },
        },
        curator: {
          connect: {
            id: data["curator_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.group.created`,
      Buffer.from(JSON.stringify(response)),
    );
    ch.ack(msg);
    console.log("group was created");
  } catch {
    console.log("smth went wrong");
  }
};
