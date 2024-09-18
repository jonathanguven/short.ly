# Shrink - URL Shortener

Shortly is a simple, yet powerful URL shortening service built with **Next.js**, **Golang**, and **PostgreSQL**. It allows users to create shortened URLs that redirect to longer, original URLs. Users can manage their URLs, view analytics, and more.

## Features

- **URL Shortening**: Create short URLs for any original long URLs.
- **Custom Aliases**: Logged-in users can create custom aliases for their shortened URLs.
- **URL Redirection**: When a user accesses the short URL, they are automatically redirected to the original URL.
- **URL Management**: Logged-in users can modify/delete their previously created shortened URLs.
- **Analytics**: Logged-in users can view how many times a shortened URL has been clicked.

## Tech Stack

- **Frontend**: React, Shadcn
- **Backend**: Golang
- **Database**: PostgreSQL
- **Monitoring**: Prometheus & Grafana

## How it Works

1. **URL Shortening**:
   - Users can submit a long URL and receive a shortened version.
   - Guest users' URLs expire after 7 days, while registered users' URLs don't expire and can be managed.
2. **Redirection**:
   - When a user visits the shortened URL, the system retrieves the original URL from the database and redirects the user to it.
3. **Analytics**:
   - Each time a shortened URL is accessed, the click count is incremented, and this data can be viewed by the URL's creator.
