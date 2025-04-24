-- Modify "movie_quotes" table
ALTER TABLE "movie_quotes" ADD COLUMN "language_quotes" bigint NULL, ADD CONSTRAINT "movie_quotes_languages_quotes" FOREIGN KEY ("language_quotes") REFERENCES "languages" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Drop "movie_quote_language" table
DROP TABLE "movie_quote_language";
