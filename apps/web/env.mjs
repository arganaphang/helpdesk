import { createEnv } from "@t3-oss/env-nextjs";
import { z } from "zod";

export const env = createEnv({
  server: {
    BASE_API_URL: z.string(),
  },
  client: {},
  runtimeEnv: {
    BASE_API_URL: process.env.BASE_API_URL,
  },
});
