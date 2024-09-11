/** @type {import('next').NextConfig} */
const nextConfig = {
  async redirects() {
    return [
      {
        source: '/s/:slug*',
        destination: 'https://api.shrink.lol/s/:slug*',
        permanent: true,
      },
    ];
  },
};

export default nextConfig;
