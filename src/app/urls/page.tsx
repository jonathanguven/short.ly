import { cookies } from "next/headers";
import { validateToken } from "@/lib/auth";
import { redirect } from "next/navigation";
import { getURLs } from "@/lib/getURLs"
import URLsClient from "./urls";

export default async function URLs() {
  const token = cookies().get('token')?.value;
  const { isAuthenticated } = validateToken(token);

  if (!isAuthenticated) {
    redirect('/login')
  }

  const initialUrls = await getURLs();

  return (
    <div className="container mx-auto py-10 lg:max-w-5xl">
      <URLsClient initialUrls={initialUrls} />
    </div>
  );
}