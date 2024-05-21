import amqplib from "amqplib/callback_api";
import { amqpConfig } from "../../../config";
import { prisma } from "../../../prisma";

export const MainEmployeeCreated = async (
  ch: amqplib.Channel,
  msg: amqplib.Message,
) => {
  try {
    const employeeData = JSON.parse(msg.content.toString());
    const employee = await prisma.user.create({
      data: {
        id: employeeData["id"],
        firstname: employeeData["firstname"],
        middlename: employeeData["middlename"],
        surname: employeeData["surname"],
        email: employeeData["email"],
        roleType: employeeData["role_type"],
        employee: {
          create: {
            id: employeeData["role_id"],
          },
        },
      },
    });
    ch.publish(
      amqpConfig.exchange,
      `${amqpConfig.queue_name}.employee.created`,
      Buffer.from(JSON.stringify(employee)),
    );
    console.log("employee was created");
    ch.ack(msg);
  } catch {
    console.log("smth went wrong");
  }
};
