package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"api/ent"
	"api/ent/character"
	"api/ent/movie"
)

type SeedQuote struct {
	Movie     string `json:"movie"`
	Year      int    `json:"year"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func Seed(ctx context.Context, client *ent.Client) error {
	log.Println("🌱 Checking if database needs seeding...")

	count, err := client.MovieQuote.Query().Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("✅ Database already seeded. Skipping.")
		return nil
	}

	file, err := os.Open("api/db/movie_quotes_seed.json")
	if err != nil {
		return fmt.Errorf("failed to open seed file: %w", err)
	}
	defer file.Close()

	var quotes []SeedQuote
	if err := json.NewDecoder(file).Decode(&quotes); err != nil {
		return fmt.Errorf("failed to decode seed file: %w", err)
	}

	log.Printf("📦 Seeding %d quotes...\n", len(quotes))

	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	movieCache := make(map[string]*ent.Movie)
	characterCache := make(map[string]*ent.Character)

	const batchSize = 20
	for i, q := range quotes {
		mv, ok := movieCache[q.Movie]
		if !ok {
			mv, err = tx.Movie.
				Query().
				Where(movie.TitleEQ(q.Movie)).
				Only(ctx)
			if ent.IsNotFound(err) {
				mv, err = tx.Movie.
					Create().
					SetTitle(q.Movie).
					SetYear(q.Year).
					Save(ctx)
			}
			if err != nil {
				return fmt.Errorf("movie error: %w", err)
			}
			movieCache[q.Movie] = mv
		}

		char, ok := characterCache[q.Character]
		if !ok {
			char, err = tx.Character.
				Query().
				Where(character.NameEQ(q.Character)).
				Only(ctx)
			if ent.IsNotFound(err) {
				char, err = tx.Character.
					Create().
					SetName(q.Character).
					Save(ctx)
			}
			if err != nil {
				return fmt.Errorf("character error: %w", err)
			}
			characterCache[q.Character] = char
		}

		_, err = tx.MovieQuote.
			Create().
			SetQuote(q.Quote).
			SetMovie(mv).
			SetCharacter(char).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("quote error: %w", err)
		}

		if (i+1)%batchSize == 0 {
			log.Printf("📈 Inserted %d/%d quotes...", i+1, len(quotes))
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("✅ Seeding complete.")
	return nil
}
