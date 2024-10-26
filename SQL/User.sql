CREATE TABLE "users" (
  "user_id" integer PRIMARY KEY,
  "name" varchar(255),
  "email" varchar(255),
  "password" varchar(255),
  "phone_number" varchar(20),
  "role" varchar(50),
  "created_at" timestamp
);

CREATE TABLE "otps" (
  "otp_id" integer PRIMARY KEY,
  "user_id" integer,
  "otp_code" varchar(10),
  "created_at" timestamp,
  "expires_at" timestamp,
  "status" varchar(50)
);

COMMENT ON COLUMN "users"."role" IS 'User roles, e.g., admin, user, guest';

COMMENT ON COLUMN "otps"."otp_code" IS 'OTP code sent to the user';

COMMENT ON COLUMN "otps"."status" IS 'OTP status: pending, used, expired';

ALTER TABLE "otps" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
