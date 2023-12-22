export async function GET(request: Request, { params }: { params: { path: string[] } }) {
	return Response.json({ message: "OK", url: params.path.join("/") });
}

export async function HEAD(request: Request, { params }: { params: { path: string[] } }) {}

export async function POST(request: Request, { params }: { params: { path: string[] } }) {}

export async function PUT(request: Request, { params }: { params: { path: string[] } }) {}

export async function DELETE(request: Request, { params }: { params: { path: string[] } }) {}

export async function PATCH(request: Request, { params }: { params: { path: string[] } }) {}
