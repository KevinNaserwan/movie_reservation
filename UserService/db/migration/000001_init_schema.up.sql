CREATE TABLE
    "users" (
        "id" uuid PRIMARY KEY,
        "username" varchar NOT NULL,
        "email" varchar NOT NULL,
        "role" varchar NOT NULL,
        "otp" integer,
        "verified_at" timestamp,
        "created_at" timestamp,
        "updated_at" timestamp
    );