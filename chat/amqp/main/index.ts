import amqplib from "amqplib/callback_api";
import {
  MainFacultyCreated,
  MainFacultyDeleted,
  MainFacultyUpdated,
} from "./faculty";
import {
  MainTeacherCreated,
  MainTeacherUpdated,
  MainTeacherDeleted,
} from "./teacher";
import {
  MainGroupCreated,
  MainGroupCuratorUpdated,
  MainGroupDeleted,
} from "./group";
import { MainStudentCreated, MainStudentDeleted } from "./student";
import { MainDepartmentCreated, MainDepartmentDeleted } from "./department";
import { MainEmployeeCreated } from "./employee/created";

export default async function MainRouter(
  ch: amqplib.Channel,
  msg: amqplib.Message,
) {
  switch (msg.fields.routingKey) {
    case "main.faculty.created":
      MainFacultyCreated(ch, msg);
      break;
    case "main.faculty.updated":
      MainFacultyUpdated(ch, msg);
      break;
    case "main.faculty.deleted":
      MainFacultyDeleted(ch, msg);
      break;
    case "main.department.created":
      MainDepartmentCreated(ch, msg);
      break;
    case "main.department.deleted":
      MainDepartmentDeleted(ch, msg);
      break;
    case "main.group.created":
      MainGroupCreated(ch, msg);
      break;
    case "main.group.curator_updated":
      MainGroupCuratorUpdated(ch, msg);
      break;
    case "main.group.created":
      MainGroupDeleted(ch, msg);
      break;
    case "main.employee.created":
      MainEmployeeCreated(ch, msg);
      break;
    case "main.teacher.created":
      MainTeacherCreated(ch, msg);
      break;
    case "main.teacher.updated":
      MainTeacherUpdated(ch, msg);
      break;
    case "main.teacher.deleted":
      MainTeacherDeleted(ch, msg);
      break;
    case "main.student.created":
      MainStudentCreated(ch, msg);
      break;
    case "main.student.deleted":
      MainStudentDeleted(ch, msg);
      break;
    default:
      console.log(msg.fields.routingKey, msg.content.toString());
      ch.ack(msg);
  }
}
