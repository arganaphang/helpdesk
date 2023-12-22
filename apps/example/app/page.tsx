"use client";
import React from "react";

export default function Page() {
	const ref = React.useRef<HTMLFormElement>(null);
	const [isLoading, setIsLoading] = React.useState(false);

	async function onSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setIsLoading(true);
		try {
			const formData = new FormData(event.currentTarget);
			const response = await fetch("/api/submit", {
				method: "POST",
				body: formData,
			});
			if (response.status === 201) {
				ref.current?.reset();
			}
		} catch (e) {
		} finally {
			setIsLoading(false);
		}
	}

	return (
		<div className="w-full min-h-screen flex justify-center items-center">
			<form onSubmit={onSubmit} className="flex flex-col gap-2 w-96" ref={ref}>
				<div>
					<label htmlFor="name" className="block text-xs font-medium leading-6 text-gray-500">
						Name
					</label>
					<input
						id="name"
						name="name"
						type="text"
						className="mt-1 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6"
					/>
				</div>
				<div>
					<label htmlFor="email" className="block text-xs font-medium leading-6 text-gray-500">
						Email address
					</label>
					<input
						id="email"
						name="email"
						type="email"
						className="mt-1 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6"
					/>
				</div>
				<div>
					<label htmlFor="title" className="block text-xs font-medium leading-6 text-gray-500">
						Title
					</label>
					<input
						id="title"
						name="title"
						type="text"
						className="mt-1 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6"
					/>
				</div>
				<div>
					<label htmlFor="detail" className="block text-xs font-medium leading-6 text-gray-500">
						Detail
					</label>
					<textarea
						id="detail"
						name="detail"
						rows={5}
						className="mt-1 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6"
					/>
				</div>
				<div className="mt-6 flex items-center justify-end gap-x-6">
					<button
						disabled={isLoading}
						type="submit"
						className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
					>
						Submit
					</button>
				</div>
			</form>
		</div>
	);
}
