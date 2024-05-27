import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainGroupCuratorUpdated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const groupCuratorUpdatedData = JSON.parse(msg.content.toString());
    const groupCuratorUpdated = await prisma.group.update({
      where: {
        id: groupCuratorUpdatedData["id"],
      },
      data: {
        curator: {
          connect: {
            id: groupCuratorUpdatedData["curator_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.group.curator_updated`,
      Buffer.from(JSON.stringify(groupCuratorUpdated)),
    );
    ch.ack(msg as amqplib.Message);
    console.log("group was updated");
  } catch {
    console.log("smth went wrong");
  }
};
