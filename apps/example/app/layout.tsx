import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./tailwind.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
	title: "Example Website",
	description: "Example Website to submit the issue",
};

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}): JSX.Element {
	return (
		<html lang="en">
			<body className={inter.className}>{children}</body>
		</html>
	);
}
