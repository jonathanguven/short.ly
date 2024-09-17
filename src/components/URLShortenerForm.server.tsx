import { cookies } from "next/headers";
import { validateToken } from "@/lib/auth";
import URLShortenerClient from "@/components/URLShortenerClient";

export default function URLShortenerForm() {
  const token = cookies().get('token')?.value;
  const { isAuthenticated } = validateToken(token);

  return <URLShortenerClient isAuthenticated={isAuthenticated} />;
}
