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
		WithCharacter().
		Limit(1).
		Order(sql.OrderByRand()).           // Randomize the order
		First(ctx)
	if err != nil {
		return nil, err
	}

	return &quotespb.Quote{
		Quote:     q.Quote,
		Movie:     q.Edges.Movie.Title,
		Character: q.Edges.Character.Name,
		Year:      int32(q.Edges.Movie.Year),
	}, nil
}

func (s *quoteServer) GetQuotes(ctx context.Context, req *quotespb.QuoteRequest) (*quotespb.QuoteList, error) {
	query := s.client.MovieQuote.Query().
		WithMovie().
		WithCharacter()

	quotes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &quotespb.QuoteList{}
	for _, q := range quotes {
		result.Quotes = append(result.Quotes, &quotespb.Quote{
			Quote:     q.Quote,
			Movie:     q.Edges.Movie.Title,
			Character: q.Edges.Character.Name,
			Year:      int32(q.Edges.Movie.Year),
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
