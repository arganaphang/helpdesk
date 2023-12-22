"use client";

import { useIssues } from "@/services/issue";

export default function Page(): JSX.Element {
	const { data } = useIssues();
	return (
		<>
			<h1>Hello from Dashboard Website</h1>
			<pre>{JSON.stringify(data, null, 2)}</pre>
		</>
	);
}
