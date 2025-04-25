package grpc

import (
	"context"
	"log"
	"net"

	"api/ent"
	quotespb "api/proto/quotespb"

	"entgo.io/ent/dialect/sql"
	"google.golang.org/grpc"
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
		WithMovie().
		WithLanguage().
		Limit(1).
		Order(sql.OrderByRand()).           // Randomize the order
		First(ctx)
	if err != nil {
		return nil, err
	}

	category, err := s.client.Movie.QueryCategory(q.Edges.Movie).First(ctx)
	if err != nil {
		return nil, err
	}

	character, err := s.client.Movie.QueryCharacters(q.Edges.Movie).First(ctx)
	if err != nil {
		println(err)
		return nil, err
	}


	return &quotespb.Quote{
		Quote:     	q.Quote,
		Movie:    	&quotespb.Movie{
			Name:		q.Edges.Movie.Title,
			Character:  &quotespb.Character{
				Name: 	character.Name,
				Actor: 	character.Actor,
			},
			Category: 	category.Name,
			Year:		int32(q.Edges.Movie.Year),
		},
		Context: 	q.Context,
		Language: 	q.Edges.Language.Name,
	}, nil
}

func (s *quoteServer) GetQuotes(ctx context.Context, req *quotespb.QuoteRequest) (*quotespb.QuoteList, error) {
	quotes, err := s.client.MovieQuote.Query().
		WithMovie(func(q *ent.MovieQuery) {
			q.WithCharacters()
			q.WithCategory()
		}).
		WithLanguage().
		All(ctx)
	
	if err != nil {
		return nil, err
	}

	result := &quotespb.QuoteList{}
	for _, q := range quotes {
		result.Quotes = append(result.Quotes, &quotespb.Quote{
			Quote:     q.Quote,
			Movie:     &quotespb.Movie{
				Name:		q.Edges.Movie.Title,
				Character:  &quotespb.Character{
					Name: 	q.Edges.Movie.Edges.Characters[0].Name,
					Actor: 	q.Edges.Movie.Edges.Characters[0].Actor,
				},
				Category: 	q.Edges.Movie.Edges.Category[0].Name,
				Year:		int32(q.Edges.Movie.Year),
			},
			Context: 	q.Context,
			Language:  q.Edges.Language.Name,
		})
	}

	return result, nil
}

func StartGRPC(client *ent.Client) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	quotespb.RegisterQuoteServiceServer(s, NewQuoteServer(client))

	log.Println("🚀 gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
