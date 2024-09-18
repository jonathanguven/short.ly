import { URL } from "@/app/urls/columns";
import { cookies } from "next/headers";

export async function getURLs(): Promise<URL[]> {
  try {
    const token = cookies().get('token')?.value;

    if (!token) {
      throw new Error("No token available");
    }

    const response = await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}/urls`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Cookie': `token=${token}`
      }
    });

    if (!response.ok) {
      throw new Error(response.statusText);
    }

    const urls: URL[] = await response.json();
    return urls;
  } catch (error) {
    console.error(error);
    throw new Error('Failed to fetch URLs');
  }
}