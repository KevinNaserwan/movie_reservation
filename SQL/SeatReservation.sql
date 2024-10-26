CREATE TABLE "seat_reservations" (
  "reservation_id" integer PRIMARY KEY,
  "show_id" integer,
  "seat_number" varchar(10),
  "user_id" integer,
  "status" varchar(50),
  "reserved_at" timestamp
);

CREATE TABLE "showtimes" (
  "show_id" integer PRIMARY KEY,
  "movie_id" integer,
  "theater_id" integer,
  "start_time" timestamp,
  "end_time" timestamp
);

ALTER TABLE "seat_reservations" ADD FOREIGN KEY ("show_id") REFERENCES "showtimes" ("show_id");
