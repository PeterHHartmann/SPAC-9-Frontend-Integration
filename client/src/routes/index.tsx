import { createFileRoute } from '@tanstack/react-router';
import { createClient } from '@connectrpc/connect';
import { createGrpcWebTransport } from '@connectrpc/connect-web';
import { useEffect, useState } from 'react';
import logo from '../logo.svg';
import type { Transport } from '@connectrpc/connect';
import type { Quote } from '@/proto/gen/quotes_pb';
import { QuoteService } from '@/proto/gen/quotes_pb';

export const Route = createFileRoute('/')({
	component: App,
});

const apiUrl = "http://localhost:8000";

const transport: Transport = createGrpcWebTransport({
	baseUrl: apiUrl
});

const quoteClient = createClient(QuoteService, transport);

const getQuotes = async () => {
	const response = await quoteClient.getQuotes({});
	console.log(response.quotes);

	return response.quotes;
};

function App() {
	const [quotes, setQuotes] = useState<Array<Quote>>();

	useEffect(() => {
		getQuotes().
			then((q) => setQuotes(q));
	}, []);

	console.log(quotes);

	return (
		<div className="text-center">
			<header className="min-h-screen flex flex-col items-center justify-center bg-[#282c34] text-white text-[calc(10px+2vmin)]">

				<img
					src={logo}
					className="h-[40vmin] pointer-events-none animate-[spin_20s_linear_infinite]"
					alt="logo"
				/>
				<p>
					Edit <code>src/routes/index.tsx</code> and save to reload.
				</p>
				<a
					className="text-[#61dafb] hover:underline"
					href="https://reactjs.org"
					target="_blank"
					rel="noopener noreferrer"
				>
					Learn React
				</a>
				<a
					className="text-[#61dafb] hover:underline"
					href="https://tanstack.com"
					target="_blank"
					rel="noopener noreferrer"
				>
					Learn TanStack
				</a>
			</header>
		</div>
	);
}
