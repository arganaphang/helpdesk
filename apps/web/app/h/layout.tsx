import type { Metadata } from "next";

export const metadata: Metadata = {
	title: "Helpdesk",
	description: "Helpdesk",
};

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}): JSX.Element {
	return <>{children}</>;
}
