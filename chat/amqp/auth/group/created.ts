import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const AuthGroupCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const groupData = JSON.parse(msg.content.toString());
    console.log(groupData);
    const group = await prisma.group.create({
      data: {
        id: groupData["id"],
        name: groupData["name"],
        faculty: {
          connect: {
            id: groupData["faculty_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.group.created`,
      Buffer.from(JSON.stringify(group)),
    );
    ch.ack(msg);
    console.log("group was created");
  } catch {
    console.log("smth went wrong");
  }
};
