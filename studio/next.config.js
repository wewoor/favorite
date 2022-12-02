/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  swcMinify: true,

  async rewrites() {
    return [{
      source: `${process.env.REMOTE_PREFIX}/:path*`,
      destination: `${process.env.REMOTE_HOST}/:path*`
    }]
  }
}

module.exports = nextConfig
