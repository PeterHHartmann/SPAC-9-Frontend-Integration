import { useQuoteRandom } from '@/view-models/quote-view-models';
import type React from 'react';

export const RandomQuoteView: React.FC = () => {
	const { data, isLoading, error } = useQuoteRandom();

	if (isLoading) {
		return <div className="p-6">Loading quote...</div>;
	}

	if (error) {
		return <div className="p-6 text-red-500">Error: {error.message}</div>;
	}

	return (
		<div className="w-max mx-auto">
			<div className='flex flex-col p-8 text-2xl bg-slate-100 rounded-2xl border-1'>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Id:</p>
					<p>{data?.quote.id} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Quote:</p>
					<p>{data?.quote.quote} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Character:</p>
					<p>{data?.quote.character?.name} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Movie:</p>
					<p>{data?.quote.movie?.title} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Year:</p>
					<p>{data?.quote.movie?.year} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Genre:</p>
					<p>{data?.quote.movie?.category} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Actor:</p>
					<p>{data?.quote.character?.actor} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Context:</p>
					<p>{data?.quote.context} </p>
				</div>
				<div className='grid grid-cols-2 mb-2 pb-2 px-2 border-b-1'>
					<p>Language:</p>
					<p>{data?.quote.language} </p>
				</div>
			</div>
		</div>
	);
};