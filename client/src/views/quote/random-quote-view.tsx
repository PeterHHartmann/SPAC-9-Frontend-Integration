import { queryClient, queryKeys } from '@/api/queryClient';
import { Button } from '@/components/ui/button';
import { useQuoteRandom } from '@/view-models/quote-view-models';
import { RefreshCw } from 'lucide-react';
import type React from 'react';

export const RandomQuoteView: React.FC = () => {
	const { data, isLoading, error } = useQuoteRandom();

	const refreshRandomQuote = () => {
		queryClient.invalidateQueries({ queryKey: [...queryKeys.quotes.random] });
	};

	if (isLoading) {
		return <div className="p-6">Loading quote...</div>;
	}

	if (error) {
		return <div className="p-6 text-red-500">Error: {error.message}</div>;
	}

	return (
		<div className='grid grid-min-rows h-max gap-4'>
			<div className="max-w-xl h-min mx-auto bg-white dark:bg-gray-800 shadow-xl rounded-2xl p-6 transition hover:scale-[1.02] duration-200">
				<blockquote className="text-xl italic text-gray-800 dark:text-gray-100 border-l-4 border-blue-500 pl-4 mb-4">
					“{data?.quote.quote}”
				</blockquote>

				<div className="text-sm text-gray-600 dark:text-gray-300 mb-6">
					— {data?.quote.character?.name} in
					<span className="font-semibold"> {data?.quote.movie?.title} </span>
					<span> ({data?.quote.movie?.year}), </span>
					<span>played by {data?.quote.character?.actor}</span>
				</div>

				<div className="grid grid-cols-2 gap-4 text-sm text-gray-700 dark:text-gray-400">
					<div>
						<span className="font-medium">Genre:</span> {data?.quote.movie?.category}
					</div>
					<div>
						<span className="font-medium">Language:</span> {data?.quote.language}
					</div>
					<div className="col-span-2">
						<span className="font-medium">Context:</span>{" "}
						<span className="italic">{data?.quote.context}</span>
					</div>
				</div>
			</div>
			<div className='flex w-full justify-center' >
				<Button size="lg" onClick={() => refreshRandomQuote()} >
					<RefreshCw />
					Get another quote
				</Button>
			</div>
		</div>
	);
};