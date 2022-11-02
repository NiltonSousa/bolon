import fastify from "fastify";
import cors from "@fastify/cors";
import { PrismaClient } from "@prisma/client";
import { z } from "zod";
import ShortUniqueId from "short-unique-id";

const prisma = new PrismaClient({
  log: ["query"],
});

async function bootstrap() {
  const app = fastify({
    logger: true,
  });

  await app.register(cors, {
    origin: true,
  });

  app.get("/pools/count", async () => {
    const count = await prisma.pool.count();

    return { count };
  });

  app.get("/users/count", async () => {
    const count = await prisma.user.count();

    return { count };
  });

  app.get("/guesses/count", async () => {
    const count = await prisma.guess.count();

    return { count };
  });

  app.post("/pools", async (request, reply) => {
    const createPoolBody = z.object({
      title: z.string(),
    });
    const { title } = createPoolBody.parse(request.body);
    const generateId = new ShortUniqueId({ length: 6 });
    const code = String(generateId()).toUpperCase();

    await prisma.pool.create({
      data: {
        title,
        code,
      },
    });

    return reply.status(201).send({ code });
  });

  await app.listen({ port: 8080, host: "0.0.0.0" });
}

bootstrap();
