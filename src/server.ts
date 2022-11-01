import fastify from "fastify";

async function bootstrap() {
  const app = fastify({
    logger: true,
  });

  app.get("/pools/count", () => {
    return { count: 0 };
  });

  await app.listen({ port: 8080 });
}

bootstrap();
