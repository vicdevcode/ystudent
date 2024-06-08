import { JwtPayload } from "jsonwebtoken";
import { InstanceOfJwt, JwtUser, getPayload } from "./jwt";
import { Socket } from "socket.io";
import { prisma } from "./prisma";

type UserPayload = JwtUser | JwtPayload | string;

export const socketImp = (socket: Socket) => {
  console.log("a user connected", socket.id);
  const access_token = socket.handshake.auth.access_token;
  const payload: UserPayload = getPayload(access_token);

  socket.on("send_message_to", async (data) => {
    let dataJwt = {
      email: "",
    };
    if (InstanceOfJwt(payload)) dataJwt.email = payload["email"];
    const user = await prisma.user.findUnique({
      where: {
        email: dataJwt.email,
      },
    });
    if (!user?.id) return;
    const newMessage = await prisma.message.create({
      data: {
        chatId: data.chat_id,
        content: data.message,
        senderId: user?.id,
        important: false,
        senderFio: user.surname + " " + user.firstname + " " + user.middlename,
      },
    });
    socket.to(data.chat_id).emit("receive_message", {
      senderId: newMessage.senderId,
      senderFio: newMessage.senderFio,
      content: data.message,
      createdAt: newMessage.createdAt,
    });
  });

  socket.on("join_room", (chat_id) => {
    socket.join(chat_id);
    socket.on("send_news", async (message) => {
      let dataJwt = {
        email: "",
      };
      if (InstanceOfJwt(payload)) dataJwt.email = payload["email"];
      const user = await prisma.user.findUnique({
        where: {
          email: dataJwt.email,
        },
      });
      if (!user?.id) return;

      const chat = await prisma.chat.findUnique({
        where: {
          id: chat_id,
        },
      });
      if (chat?.type != "NEWS") return;

      console.log(chat.id, chat.name)

      const allUsers = await prisma.user.findMany({
        where: {
          chats: {
            some: {
              id: chat.id,
            },
          },
        },
        include: {
          chats: true,
        },
      });

      await prisma.message.createMany({
        data: [
          {
            chatId: chat_id,
            content: message,
            senderId: user?.id,
            is_news: true,
            senderFio:
              user.surname + " " + user.firstname + " " + user.middlename,
          },
          ...allUsers.map((u) => {
            const userNewsChat = u.chats.find(
              (a) => a.type === "USER_NEWS",
            );
            if (userNewsChat) 
            return {
              chatId: userNewsChat.id,
              content: message,
              senderId: user?.id,
              is_news: true,
              senderFio:
                user.surname + " " + user.firstname + " " + user.middlename,
            };
            return {
              chatId: chat_id,
              content: message,
              senderId: user?.id,
              is_news: true,
              senderFio:
                user.surname + " " + user.firstname + " " + user.middlename,
            };
          }),
        ],
      });
      socket.to(chat_id).emit("receive_message", {
        senderId: user.id,
        senderFio: user.surname + " " + user.firstname + " " + user.middlename,
        content: message,
        createdAt: new Date().toISOString(),
      });
      for (let i = 0; i < allUsers.length; i++) {
        const userNewsChat = allUsers[i].chats.find(
          (a) => a.type === "USER_NEWS",
        );
        if (userNewsChat) 
        socket.to(userNewsChat.id).emit("receive_message", {
          senderId: user.id,
          senderFio:
            user.surname + " " + user.firstname + " " + user.middlename,
          content: message,
          createdAt: new Date().toISOString(),
        });
      }
    });
    socket.on("send_message", async (message) => {
      let data = {
        email: "",
      };
      if (InstanceOfJwt(payload)) data.email = payload["email"];
      const user = await prisma.user.findUnique({
        where: {
          email: data.email,
        },
      });
      if (!user?.id) return;
      const newMessage = await prisma.message.create({
        data: {
          chatId: chat_id,
          content: message,
          senderId: user?.id,
          senderFio:
            user.surname + " " + user.firstname + " " + user.middlename,
        },
      });
      socket.to(chat_id).emit("receive_message", {
        senderId: newMessage.senderId,
        senderFio: newMessage.senderFio,
        content: message,
        createdAt: newMessage.createdAt,
      });
    });
    socket.on("leave_room", () => {
      socket.leave(chat_id);
    });
  });
  socket.on("disconnect", (reason) => {
    console.log(`[${socket.id}] socket disconnected - ${reason}`);
  });
};
