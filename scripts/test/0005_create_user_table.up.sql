CREATE TABLE IF NOT EXISTS "user" (
                                       "id" SERIAL PRIMARY KEY,
                                       "name" VARCHAR(255) NOT NULL,
                                       "email" VARCHAR(255)  NOT NULL,
                                       "password" TEXT NOT NULL,
                                       "role" VARCHAR(255) NOT NULL
);