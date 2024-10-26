CREATE TABLE "theaters" (
  "theater_id" integer PRIMARY KEY,
  "name" varchar(255),
  "location" varchar(255),
  "total_rooms" integer
);

CREATE TABLE "rooms" (
  "room_id" integer PRIMARY KEY,
  "theater_id" integer,
  "room_number" varchar(50),
  "capacity" integer
);

ALTER TABLE "rooms" ADD FOREIGN KEY ("theater_id") REFERENCES "theaters" ("theater_id");
