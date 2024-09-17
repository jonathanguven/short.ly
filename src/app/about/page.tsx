import { cookies } from "next/headers";

export default function About() {
  const cookieStore = cookies();
  const token = cookieStore.get('token');

  return (
    <div>
      about
    </div>
  );
}
