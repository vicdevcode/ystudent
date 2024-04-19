export const amqp_url = process.env.AMQP_URL as string;
export const amqp_exchange = process.env.AMQP_EXCHANGE as string;
export const amqp_topics = (process.env.AMQP_TOPICS as string).split(",");
export const amqp_queue_name = process.env.AMQP_QUEUE_NAME as string;
