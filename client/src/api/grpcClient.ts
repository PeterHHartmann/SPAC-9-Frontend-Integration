import { createClient } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import type { Transport } from '@connectrpc/connect';
import { QuoteService } from '@/proto/gen/quotes_pb';
import type { Quote } from '@/proto/gen/quotes_pb';

const apiUrl = 'http://localhost:8000';

export const grpcWebTransport: Transport = createGrpcWebTransport({
	baseUrl: apiUrl
});

export const quoteClient = createClient(QuoteService, grpcWebTransport);

export const getQuotes = async (): Promise<Array<Quote>> => {
	const response = await quoteClient.getQuotes({});
	console.log(response.quotes);

	return response.quotes;
};