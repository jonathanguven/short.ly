import { cookies } from "next/headers";
import Link from "next/link";
import { validateToken } from "@/lib/auth";
import LogoutButton from "./LogoutButton";

export const Header = () => {
  const token = cookies().get('token')?.value;
  const { isAuthenticated, user } = validateToken(token);

  return (
    <header className="bg-white dark:bg-gray-800 shadow">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex justify-between items-center">
          <Link href="/" className="text-2xl font-bold text-gray-900 dark:text-white">
            Shrink.lol
          </Link>

          <nav>
            <ul className="flex items-center space-x-4">
              <li>
                <Link href="/about" className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white">
                  About
                </Link>
              </li>
              <li className="flex place-items-center gap-4 pl-4">
                {!isAuthenticated ? (
                  <Link href="/login" className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white">
                    Login
                  </Link>
                ) : (
                  <>
                    <Link
                      href='/urls'
                      className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                    >
                      My URLs
                    </Link>
                    <LogoutButton />
                  </>
                )}
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </header>
  );
};
