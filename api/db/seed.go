package db

import (
	"api/ent"
	"api/ent/category"
	"api/ent/character"
	"api/ent/language"
	"api/ent/movie"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JsonMovieQuote struct {
	Quote     string `json:"quote"`
	Movie     string `json:"movie"`
	Character string `json:"character"`
	Actor	  string `json:"actor"`
	Category  string `json:"category"`
	Year      int    `json:"year"`
	Context   string `json:"context"`
	Language  string `json:"language"`
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

	file, err := os.Open("db/movie_quotes_seed.json")
	if err != nil {
		return fmt.Errorf("failed to open seed file: %w", err)
	}
	defer file.Close()

	var quotes []JsonMovieQuote
	if err := json.NewDecoder(file).Decode(&quotes); err != nil {
		return fmt.Errorf("failed to decode seed file: %w", err)
	}

	var bulkCreates []*ent.MovieQuoteCreate

	for _, q := range quotes {

		ct, err := client.Category.
			Query(). 
			Where(category.NameEQ(q.Category)). 
			Only(ctx) 
		if err != nil {
			ct, err = client.Category. 
				Create(). 
				SetName(q.Category). 
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create category: %v", err)
				continue
			}
		}

		// Fetch or create Movie
		mv, err := client.Movie.
			Query().
			Where(movie.TitleEQ(q.Movie)).
			Only(ctx)
		if err != nil {
			mv, err = client.Movie.
				Create().
				SetTitle(q.Movie).
				SetYear(q.Year).
				AddCategory(ct).
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create movie: %v", err)
				continue
			}
		}

		// Fetch or create Character
		_, err = client.Character.
			Query().
			Where(character.NameEQ(q.Character)).
			Only(ctx)
		if err != nil {
			_, err = client.Character.
				Create().
				SetName(q.Character).
				SetActor(q.Actor).
				SetMovie(mv).
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create character: %v", err)
				continue
			}
		}

		// Fetch or create Language
		lang, err := client.Language.
			Query().
			Where(language.NameEQ(q.Language)).
			Only(ctx)
		if err != nil {
			lang, err = client.Language.
				Create().
				SetName(q.Language).
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create Language: %v", err)
				continue
			}
		}

		// Build MovieQuote entry
		create := client.MovieQuote.
			Create().
			SetQuote(q.Quote).
			SetContext(q.Context).
			SetMovie(mv).
			SetLanguage(lang)

		bulkCreates = append(bulkCreates, create)
	}

	if _, err := client.MovieQuote.CreateBulk(bulkCreates...).Save(ctx); err != nil {
		return fmt.Errorf("failed to create movie quotes: %w", err)
	}

	log.Println("✅ Seeding complete.")
	return nil
}