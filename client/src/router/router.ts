import { RootLayout } from '@/components/layout/RootLayout';
import { AllQuotesView } from '@/views/quote/all-quotes-view';
import { RandomQuoteView } from '@/views/quote/random-quote-view';
import { createBrowserRouter, redirect } from 'react-router-dom';

export const router = createBrowserRouter([
	{
		path: "/",
		Component: RootLayout,
		children: [
			// Default redirect to /products
			{
				index: true,
				loader: async () => redirect("/quotes/random"),
			},
			// Product routes
			{ path: "quotes/random", Component: RandomQuoteView },
			{ path: "quotes", Component: AllQuotesView },
		],
	},
]);