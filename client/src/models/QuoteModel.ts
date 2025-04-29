import { createClient, type Client } from '@connectrpc/connect';
import { QuoteService, type Quote } from '@/proto/gen/quotes_pb';
import { grpcWebTransport, quoteClient } from '@/api/grpcClient';

export class QuoteModel {
	quoteClient: Client<typeof QuoteService> = createClient(QuoteService, grpcWebTransport);

	static async getAll(): Promise<{ quotes: Array<Quote>; }> {
		const response = await quoteClient.getQuotes({});
		return {
			quotes: response.quotes
		};

	}

	static async getRandom(): Promise<{ quote: Quote; }> {
		const quote = await quoteClient.getRandomQuote({});
		return {
			quote: quote
		};
	}

}