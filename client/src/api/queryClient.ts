import { QueryClient } from '@tanstack/react-query';

export const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
			gcTime: 1000 * 60 * 60 * 24, // 24 hours
			retry: 1,
			refetchOnWindowFocus: false,
		},
	},
});

export const queryKeys = {
	quotes: {
		all: ["quotes"] as const,
		random: ["random-quote"] as const,
	}
};