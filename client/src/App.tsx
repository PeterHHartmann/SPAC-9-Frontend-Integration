import { RouterProvider } from 'react-router-dom';
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from '@/api/queryClient';
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { router } from '@/router/router';

// Data router configuration


export default function App() {
	return (
		<QueryClientProvider client={queryClient}>
			<RouterProvider router={router} />
			{import.meta.env.VITE_ENABLE_QUERY_DEVTOOLS === "true" && (
				<ReactQueryDevtools initialIsOpen={false} />
			)}
		</QueryClientProvider>
	);
}

