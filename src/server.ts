import fastify from "fastify";
import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient({
  log: ["query"],
});

async function bootstrap() {
  const app = fastify({
    logger: true,
  });

  app.get("/pools/count", async () => {
    const count = await prisma.pool.count();

    return { count };
  });

  await app.listen({ port: 8080 });
}

bootstrap();
