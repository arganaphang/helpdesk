"use client";
import { QueryClient, QueryClientProvider, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

const queryClient = new QueryClient();
interface Props {
	children: React.ReactNode;
}
export default function Provider({ children }: Props) {
	return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
}
