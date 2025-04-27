package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"api/ent"
	"api/ent/category"
	"api/ent/character"
	"api/ent/language"
	"api/ent/movie"
	"api/ent/moviequote"
	"api/proto/quotespb"

	"entgo.io/ent/dialect/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type quoteServer struct {
	quotespb.UnimplementedQuoteServiceServer
	client *ent.Client
}

func NewQuoteServer(client *ent.Client) *quoteServer {
	return &quoteServer{client: client}
}

func (s *quoteServer) GetRandomQuote(ctx context.Context, _ *emptypb.Empty) (*quotespb.Quote, error) {
	// Use the SQL query to order by random in PostgreSQL
	q, err := s.client.MovieQuote.Query().
		WithMovie(func(q *ent.MovieQuery) {
			q.WithCategory()
		}).
		WithLanguage().
		WithCharacter().
		Limit(1).
		Order(sql.OrderByRand()).           // Randomize the order
		First(ctx)
	if err != nil {
		return nil, err
	}

	return &quotespb.Quote{
		Id: 	   int32(q.ID),
		Quote:     q.Quote,
		Movie:     &quotespb.Movie{
			Title:		q.Edges.Movie.Title,
			Category: 	q.Edges.Movie.Edges.Category.Name,
			Year:		int32(q.Edges.Movie.Year),
		},
		Character:  &quotespb.Character{
			Name: 	q.Edges.Character.Name,
			Actor: 	q.Edges.Character.Actor,
		},
		Context: 	q.Context,
		Language:  q.Edges.Language.Name,
	}, nil
}

func (s *quoteServer) GetQuotes(ctx context.Context, req *quotespb.QuoteRequest) (*quotespb.QuoteList, error) {
	quotes, err := s.client.MovieQuote.Query().
		WithMovie(func(q *ent.MovieQuery) {
			q.WithCategory()
		}).
		WithLanguage().
		WithCharacter().
		All(ctx)
	
	if err != nil {
		return nil, err
	}

	result := &quotespb.QuoteList{}
	for _, q := range quotes {
		result.Quotes = append(result.Quotes, &quotespb.Quote{
			Id: 	   int32(q.ID),
			Quote:     q.Quote,
			Movie:     &quotespb.Movie{
				Title:		q.Edges.Movie.Title,

				Category: 	q.Edges.Movie.Edges.Category.Name,
				Year:		int32(q.Edges.Movie.Year),
			},
			Character:  &quotespb.Character{
				Name: 	q.Edges.Character.Name,
				Actor: 	q.Edges.Character.Actor,
			},
			Context: 	q.Context,
			Language:  q.Edges.Language.Name,
		})
	}

	return result, nil
}

func (s *quoteServer) CreateQuote(ctx context.Context, req *quotespb.CreateQuoteRequest) (*quotespb.Quote, error) {

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Get or create category
	category, err := s.client.Category.
		Query(). 
		Where(category.NameEQ(req.Movie.Category)). 
		Only(ctx) 
	if err != nil {
		category, err = s.client.Category. 
			Create(). 
			SetName(req.Movie.Category). 
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Get or create language
	language, err := s.client.Language.
		Query().
		Where(language.NameEQ(req.Language)).
		Only(ctx)
	if err != nil {
		language, err = s.client.Language.
			Create().
			SetName(req.GetLanguage()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Get or create Movie
	movie, err := s.client.Movie.
		Query().
		Where(movie.TitleEQ(req.Movie.Title)).
		Only(ctx)
	if err != nil {
		movie, err = s.client.Movie.
			Create().
			SetTitle(req.Movie.Title).
			SetYear(int(req.Movie.Year)).
			SetCategory(category).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Get or create Character
	character, err := s.client.Character.
		Query().
		Where(character.NameEQ(req.Character.Name)).
		Only(ctx)
	if err != nil {
		//create Character
		character, err = s.client.Character.
			Create().
			SetName(req.Character.Name).
			SetActor(req.Character.Name).
			SetMovie(movie).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	quote, err := s.client.MovieQuote.Create().
		SetQuote(req.Quote).
		SetContext(req.Context).
		SetLanguage(language).
		SetMovie(movie).
		SetCharacter(character).
		Save(ctx)


	q, err := s.client.MovieQuote.
		Query().
		Where(moviequote.IDEQ(quote.ID)).
		Only(ctx)

	println(q.Edges.Language.ID)

	response := &quotespb.Quote{
		Id: 	  	int32(q.ID),
		Quote:     	q.Quote,
		Context: 	q.Context,
		Language:  	q.Edges.Language.Name,
		Movie:     	&quotespb.Movie{
			Title:		q.Edges.Movie.Title,
			Category: 	q.Edges.Movie.Edges.Category.Name,
			Year:		int32(q.Edges.Movie.Year),
		},
		Character:  &quotespb.Character{
			Name: 	q.Edges.Character.Name,
			Actor: 	q.Edges.Character.Actor,
		},
	}
	return response, nil


}

func StartGRPC(client *ent.Client) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	quotespb.RegisterQuoteServiceServer(s, NewQuoteServer(client))

	reflection.Register(s)

	log.Println("🚀 gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
