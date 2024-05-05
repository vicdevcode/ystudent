import amqplib from "amqplib/callback_api";
import { AuthFacultyCreated } from "./faculty";
import { AuthTeacherCreated } from "./teacher";
import { AuthGroupCreated, AuthGroupCuratorUpdated } from "./group";
import { AuthStudentCreated } from "./student";

export default async function AuthRouter(
  ch: amqplib.Channel,
  msg: amqplib.Message,
) {
  switch (msg.fields.routingKey) {
    case "auth.faculty.created":
      AuthFacultyCreated(ch, msg);
      break;
    case "auth.group.created":
      AuthGroupCreated(ch, msg);
      break;
    case "auth.group.curator_updated":
      AuthGroupCuratorUpdated(ch, msg);
      break;
    case "auth.student.created":
      AuthStudentCreated(ch, msg);
      break;
    case "auth.teacher.created":
      AuthTeacherCreated(ch, msg as amqplib.Message);
      break;
    default:
      console.log(msg.fields.routingKey, msg.content.toString());
      ch.ack(msg);
  }
}
