import { cookies } from "next/headers";
import { validateToken } from "@/lib/auth";
import { redirect } from "next/navigation";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { getURLs } from "@/lib/getURLs"


export default async function URLs() {
  const token = cookies().get('token')?.value;
  const { isAuthenticated, user } = validateToken(token);

  if (!isAuthenticated) {
    redirect('/login')
  }

  const urls = await getURLs();

  return (
    <div className="container mx-auto py-10">
      <DataTable columns={columns} data={urls} />
    </div>
  );
}
