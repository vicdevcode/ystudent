import amqplib from "amqplib/callback_api";
import { amqp_exchange, amqp_queue_name, amqp_topics } from "./constant";
import { prisma } from "../prisma";

export function start(err: Error, conn: amqplib.Connection) {
  if (err) throw err;

  conn.createChannel((err: Error, ch: amqplib.Channel) => {
    if (err) throw err;

    ch.assertExchange(amqp_exchange, "topic", {
      durable: false,
    });

    ch.assertQueue(
      amqp_queue_name,
      {
        durable: true,
      },
      (err: Error, q: amqplib.Replies.AssertQueue) => {
        if (err) throw err;

        amqp_topics.map((topic) => ch.bindQueue(q.queue, amqp_exchange, topic));

        ch.consume(
          q.queue,
          async (msg: amqplib.Message | null) => {
            const r = msg?.fields?.routingKey;
            switch (r) {
              case "auth.faculty.created":
                const facultyData = JSON.parse(
                  msg?.content.toString() as string,
                );
                const faculty = await prisma.faculty.create({
                  data: {
                    name: facultyData["name"],
                  },
                });
                ch.publish(
                  amqp_exchange,
                  `${amqp_queue_name}.faculty.created`,
                  Buffer.from(JSON.stringify(faculty)),
                );
                console.log("faculty was created");
                break;
              case "auth.group.created":
                const groupData = JSON.parse(msg?.content.toString() as string);
                const group = await prisma.group.create({
                  data: {
                    name: groupData["name"],
                    facultyId: groupData["faculty_id"],
                  },
                });
                ch.publish(
                  amqp_exchange,
                  `${amqp_queue_name}.group.created`,
                  Buffer.from(JSON.stringify(group)),
                );
                console.log("group was created");
                break;
              case "auth.group.curator_updated":
                const groupCuratorUpdatedData = JSON.parse(
                  msg?.content.toString() as string,
                );
                const groupCuratorUpdated = await prisma.group.update({
                  where: {
                    id: groupCuratorUpdatedData["id"],
                  },
                  data: {
                    curatorId: groupCuratorUpdatedData["curator_id"],
                  },
                });
                ch.publish(
                  amqp_exchange,
                  `${amqp_queue_name}.group.curator_updated`,
                  Buffer.from(JSON.stringify(groupCuratorUpdated)),
                );
                console.log("group was updated");
                break;
              case "auth.student.created":
                const studentData = JSON.parse(
                  msg?.content.toString() as string,
                );
                const user = await prisma.user.create({
                  data: {
                    firstname: studentData["firstname"],
                    middlename: studentData["middlename"],
                    surname: studentData["surname"],
                    email: studentData["email"],
                    student: {
                      create: {
                        groupId: studentData["student"]["group_id"],
                      },
                    },
                  },
                });
                ch.publish(
                  amqp_exchange,
                  `${amqp_queue_name}.student.created`,
                  Buffer.from(JSON.stringify(user)),
                );
                console.log("student was created");
                break;
              case "auth.teacher.created":
                const teacherData = JSON.parse(
                  msg?.content.toString() as string,
                );
                const teacher = await prisma.user.create({
                  data: {
                    firstname: teacherData["firstname"],
                    middlename: teacherData["middlename"],
                    surname: teacherData["surname"],
                    email: teacherData["email"],
                    teacher: {
                      create: {},
                    },
                  },
                });
                ch.publish(
                  amqp_exchange,
                  `${amqp_queue_name}.teacher.created`,
                  Buffer.from(JSON.stringify(teacher)),
                );
                console.log("teacher was created");
                break;
              default:
                console.log(r, msg?.content.toString());
            }
          },
          {
            noAck: true,
          },
        );
      },
    );
    console.log("rabbitmq started");
  });
}
