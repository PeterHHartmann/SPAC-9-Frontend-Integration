import { DataTable } from '@/components/data-table';
import type { Character, Movie, Quote } from '@/proto/gen/quotes_pb';
import { useQuotes } from '@/view-models/quote-view-models';
import type { ColumnDef } from '@tanstack/react-table';
import type React from 'react';

export const AllQuotesView: React.FC = () => {
	const { data, isLoading, error } = useQuotes();

	const columns: Array<ColumnDef<Quote | undefined, unknown>> = [
		{
			header: "Id",
			accessorKey: "id"
		},
		{
			header: "Quote",
			accessorKey: "quote"
		},
		{
			header: "Character",
			accessorKey: "character",
			cell: ({ row }) => {
				const { name }: Character = row.getValue("character");
				return name;
			}
		},
		{
			header: "Movie",
			accessorKey: "movie",
			cell: ({ row }) => {
				const { title }: Movie = row.getValue("movie");
				return title;
			}
		},
		{
			header: "Year",
			accessorKey: "movie",
			cell: ({ row }) => {
				const { year }: Movie = row.getValue("movie");
				return year;
			}
		},
		{
			header: "Genre",
			accessorKey: "movie",
			cell: ({ row }) => {
				const { category }: Movie = row.getValue("movie");
				return category;
			}
		},
		{
			header: "Actor",
			accessorKey: "character",
			cell: ({ row }) => {
				const { actor }: Character = row.getValue("character");
				return actor;
			}
		},
		{
			header: "Context",
			accessorKey: "context"
		},
		{
			header: "Language",
			accessorKey: "language"
		}
	];

	if (isLoading) {
		return <div className="p-6">Loading quote...</div>;
	}

	if (error) {
		return <div className="p-6 text-red-500">Error: {error.message}</div>;
	}

	return (
		<DataTable data={data?.quotes || []} columns={columns}>
			<h2 className='w-max text-4xl whitespace-nowrap'>
				Movie Quotes
			</h2>
		</DataTable>
	);
};