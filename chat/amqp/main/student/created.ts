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
    const department = await prisma.department.findUnique({
      where: {
        id: group.departmentId,
      },
    });
    if (!department?.chatId) throw new Error("can not get department");
    const faculty = await prisma.faculty.findUnique({
      where: {
        id: department.facultyId,
      },
    });
    if (!faculty?.chatId) throw new Error("can not get faculty");
    const response = await prisma.user.create({
      data: {
        id: data["id"],
        firstname: data["firstname"],
        middlename: data["middlename"],
        surname: data["surname"],
        email: data["email"],
        roleType: data["role"],
        chats: {
          create: [
            {
              name: "Новости",
              type: "USER_NEWS",
            },
          ],
          connect: [
            {
              id: group.chatId,
            },
            {
              id: department.chatId,
            },
            {
              id: faculty.chatId,
            },
          ],
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
