FROM node:21-alpine

WORKDIR /chat

COPY package.json package-lock.json ./

RUN npm install

COPY . ./

RUN npx prisma generate
RUN npm run build

CMD npm run start:migrate
