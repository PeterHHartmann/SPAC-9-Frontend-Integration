import { createClient } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import type { Transport } from '@connectrpc/connect';
import type { Quote } from '@/proto/gen/quotes_pb';
import { QuoteService } from '@/proto/gen/quotes_pb';

const apiUrl = "http://localhost:8000";

const transport: Transport = createGrpcWebTransport({
	baseUrl: apiUrl
});

const quoteClient = createClient(QuoteService, transport);

export const getQuotes = async (): Promise<Array<Quote>> => {
	const response = await quoteClient.getQuotes({});
	console.log(response.quotes);

	return response.quotes;
};