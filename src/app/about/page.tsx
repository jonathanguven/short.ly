import Link from "next/link";

export default function About() {
  return (
    <div className="min-h-screen bg-gray-100 dark:bg-gray-900 py-12">
      <div className="max-w-4xl mx-auto px-6 lg:px-8">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-8">
          About Shrink.lol
        </h1>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          Welcome to Shrink.lol, a modern URL shortener service built with the
          goal of providing efficient, reliable, and easy-to-use link management
          for users.
        </p>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          Project Overview
        </h2>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          Shrink.lol enables users to shorten URLs, create custom aliases, and
          track link click analytics in real time. The application allows both
          registered and guest users to generate shortened links, while
          maintaining enhanced functionality for authenticated users such as
          non-expiring links and the ability to edit or delete them.
        </p>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          Key Features
        </h2>
        <ul className="list-disc pl-6 text-lg text-gray-700 dark:text-gray-300 mb-6">
          <li>Shorten URLs with or without custom aliases.</li>
          <li>Authenticated users can manage, edit, and delete their links.</li>
          <li>Real-time click tracking and analytics.</li>
        </ul>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          Tech Stack
        </h2>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          I had been meaning to learn <strong>Go</strong> for a while now, and finally decided 
          to build this app as a means to force myself to learn the language. For my database I used <strong>PostgreSQL</strong>
          , a database known for its reliability and scalability with relational data.
        </p>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          I have always been a Svelte enthusiast, but I&apos;m always open to learning new frameworks and technologies. 
          This led me to choose <strong>React</strong> for the frontend. The application is containerized using <strong>Docker</strong> for 
          easier deployment and scaling, and hosted on <strong>Google Cloud Platform (GCP)</strong>.
        </p>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          Monitoring and observability are handled through <strong>Prometheus</strong> and <strong>Grafana</strong>
          , providing real-time insights into API usage and system performance.
        </p>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          Explore the Code
        </h2>
        <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed mb-6">
          The source code for Shrink.lol is available on GitHub. Feel free to
          explore the repository to gain insights into the architecture, code
          structure, and technologies used in the project.
        </p>

        <Link href="https://github.com/jonathanguven/shrink" className="hover:underline">
          Check out the GitHub Repository
        </Link>

        <div className="mt-10 hover:underline">
          <Link href="/">
            Back to Home
          </Link>
        </div>
      </div>
    </div>
  );
};
