package storage

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

	file, err := os.Open("storage/movie_quotes_seed.json")
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
		category := getOrCreateCategory(q, client, ctx)
		movie := getOrCreateMovie(q, category, client, ctx)
		character := getOrCreateCharacter(q, movie, client, ctx)
		language := getOrCreateLanguage(q, client, ctx)

		// Build MovieQuote entry
		create := client.MovieQuote.
			Create().
			SetQuote(q.Quote).
			SetContext(q.Context).
			SetMovie(movie).
			SetCharacter(character).
			SetLanguage(language)

		bulkCreates = append(bulkCreates, create)
	}

	if _, err := client.MovieQuote.CreateBulk(bulkCreates...).Save(ctx); err != nil {
		return fmt.Errorf("failed to create movie quotes: %w", err)
	}

	log.Println("✅ Seeding complete.")
	return nil
}

func getOrCreateCategory(q JsonMovieQuote, client *ent.Client, ctx context.Context) *ent.Category {
	category, err := client.Category.
		Query(). 
		Where(category.NameEQ(q.Category)). 
		Only(ctx) 
	if err != nil {
		category, err = client.Category. 
			Create(). 
			SetName(q.Category). 
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed to create category: %v", err)
		}
	}
	return category
}

func getOrCreateLanguage(q JsonMovieQuote, client *ent.Client, ctx context.Context) *ent.Language {
	language, err := client.Language.
		Query().
		Where(language.NameEQ(q.Language)).
		Only(ctx)
	if err != nil {
		language, err = client.Language.
			Create().
			SetName(q.Language).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed to create Language: %v", err)
		}
	}
	return language
}

func getOrCreateMovie(q JsonMovieQuote, category *ent.Category, client *ent.Client, ctx context.Context) *ent.Movie {
	movie, err := client.Movie.
		Query().
		Where(movie.TitleEQ(q.Movie)).
		Only(ctx)
	if err != nil {
		movie, err = client.Movie.
			Create().
			SetTitle(q.Movie).
			SetYear(q.Year).
			SetCategory(category).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed to create movie: %v", err)
		}
	}
	return movie
}

func getOrCreateCharacter(q JsonMovieQuote, movie *ent.Movie, client *ent.Client, ctx context.Context) *ent.Character {
	//get Character
	character, err := client.Character.
		Query().
		Where(character.NameEQ(q.Character)).
		Where(character.ActorEQ(q.Actor)).
		Only(ctx)
	if err != nil {
		//create Character
		character, err = client.Character.
			Create().
			SetName(q.Character).
			SetActor(q.Actor).
			SetMovie(movie).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed to create character: %v", err)
		}
	}
	return character
}