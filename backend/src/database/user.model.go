package database

// CREATE TABLE IF NOT EXISTS "users" ("id"   SERIAL , "uuid" UUID, "name" VARCHAR(255) NOT NULL, "email" VARCHAR(255) NOT NULL UNIQUE, "description" TEXT, "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL, "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL, PRIMARY KEY ("id"));

// CREATE UNIQUE INDEX "users_email_uuid" ON "users" ("email", "uuid")

// CREATE INDEX "users_uuid" ON "users" ("uuid")
