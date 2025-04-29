import { QuoteModel } from '@/models/QuoteModel';
import type { Quote } from '@/proto/gen/quotes_pb';
import { useQuery } from '@tanstack/react-query';

const queryKeys = {
	all: ["quotes"] as const,
	random: ["random-quote"] as const,
};

export const useQuotes = () => {
	return useQuery<{
		quotes: Array<Quote>;
	}, Error>({
		queryKey: [...queryKeys.all],
		queryFn: () => QuoteModel.getAll(),
	});
};

export const useQuoteRandom = () => {
	return useQuery<{
		quote: Quote;
	}, Error>({
		queryKey: [...queryKeys.random],
		queryFn: () => QuoteModel.getRandom(),
	});
};