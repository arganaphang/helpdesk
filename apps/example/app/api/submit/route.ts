export async function POST(request: Request) {
	const body = await request.formData();
	const data = {
		customer_name: body.get("name"),
		customer_email: body.get("email"),
		title: body.get("title"),
		detail: body.get("detail"),
	};
	const response = await fetch(`${process.env.HELPDESK_HOST}/issue`, {
		method: "POST",
		body: JSON.stringify(data),
		headers: {
			"Content-Type": "application/json",
		},
	});

	if (response.status !== 201) {
		return Response.json(
			{ message: "NOT OK" },
			{
				status: response.status,
			},
		);
	}
	return Response.json(
		{ message: "OK" },
		{
			status: response.status,
		},
	);
}
