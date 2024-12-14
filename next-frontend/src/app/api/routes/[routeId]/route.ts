//routes

import { NextRequest, NextResponse } from "next/server";

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ routeId: string }> }
) {
  const { routeId } = await params;

  const apiUrl = process.env.NEST_API_URL;

  const response = await fetch(`${apiUrl}/routes/${routeId}`, {
    cache: "force-cache",
    next: {
      tags: [`routes-${routeId}`, "routes"],
    },
  });
  const data = await response.json();
  return NextResponse.json(data);
}