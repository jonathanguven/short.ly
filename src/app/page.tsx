import URLShortenerForm from "@/components/URLShortenerForm";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col bg-gray-100" >
      <Header />
      <main className="flex-grow flex items-center justify-center px-4 sm:px-6 lg:px-8">
        <URLShortenerForm />
      </main>
      <Footer />
    </div>
  );
}
