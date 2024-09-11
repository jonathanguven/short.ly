'use client';

import { useAuth } from "@/hooks/useAuth";

export const Header = () => {
  const { isAuthenticated, user } = useAuth();

  return (
    <header className="bg-white dark:bg-gray-800 shadow">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex justify-between items-center">
          <a href="/" className="text-2xl font-bold text-gray-900 dark:text-white">Shrink.lol</a>
          <nav>
            <ul className="flex space-x-4">
              <li>
                <a href="/about" className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white">
                  About
                </a>
              </li>
              <li>
                {(!isAuthenticated) ? (
                  <a href="/login" className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white">
                    Login
                  </a>
                ) : (
                  <a
                    href={`/urls/${user}`}
                    className="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
                  >
                    My URLs
                  </a>
                )}
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </header>
  );
};
