export interface AMQP_CONFIG {
  url: string;
  exchange: string;
  topics: string[];
  queue_name: string;
}

export const amqpConfig: AMQP_CONFIG = {
  url: process.env.AMQP_URL as string,
  exchange: process.env.AMQP_EXCHANGE as string,
  topics: (process.env.AMQP_TOPICS as string).split(","),
  queue_name: process.env.AMQP_QUEUE_NAME as string,
};

export const auth_check = process.env.AUTH_CHECK as string;
export const http_port = process.env.HTTP_PORT as string;
