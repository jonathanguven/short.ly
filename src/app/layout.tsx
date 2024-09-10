import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Shrink",
  description: "Shrink is a fast, free, and easy-to-use URL shortener service that lets you shorten long links. Create custom short URLs for branding or tracking purposes.",
  keywords: "URL shortener, custom short links, shrink URL, shorten URLs, link management, free URL shortener, link tracking",
  openGraph: {
    title: "Shrink.lol - Free and Custom URL Shortener Service",
    description: "Shorten your long links with Shrink.lol and create custom short URLs for branding. Free, fast, and easy-to-use URL shortening service.",
    url: "https://shrink.lol",
    siteName: "Shrink.lol",
    // images: [
    //   {
    //     url: "/path-to-image/og-image.jpg", // make sure to add an actual image here
    //     width: 1200,
    //     height: 630,
    //     alt: "Shrink.lol - Free and Custom URL Shortener",
    //   },
    // ],
    locale: "en_US",
    type: "website",
  },
  viewport: "width=device-width, initial-scale=1.0",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
