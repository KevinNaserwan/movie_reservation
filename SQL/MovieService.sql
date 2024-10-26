CREATE TABLE "movies" (
  "movie_id" integer PRIMARY KEY,
  "title" varchar,
  "genre" varchar(100),
  "duration" integer,
  "release_date" date,
  "rating" decimal(2,1),
  "description" text
);
