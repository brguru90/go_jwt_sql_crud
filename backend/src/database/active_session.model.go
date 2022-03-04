package database

//  CREATE TABLE IF NOT EXISTS "active_sessions" ("id"   SERIAL , "uuid" UUID, "user_uuid" VARCHAR(255), "token_id" VARCHAR(255) NOT NULL UNIQUE, "ua" TEXT, "ip" VARCHAR(255), "exp" BIGINT, "status" VARCHAR(255), "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL, "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL, PRIMARY KEY ("id"));

// CREATE UNIQUE INDEX "active_sessions_uuid_token_id" ON "active_sessions" ("uuid", "token_id")

// CREATE INDEX "active_sessions_user_uuid" ON "active_sessions" ("user_uuid")

// CREATE INDEX "active_sessions_token_id" ON "active_sessions" ("token_id")
