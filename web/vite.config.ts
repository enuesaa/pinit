import { unstable_vitePlugin as remix } from '@remix-run/dev'
import { defineConfig } from 'vite'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import path from 'node:path'

export default defineConfig({
  ssr: {
    noExternal: ['@radix-ui/themes'],
  },
  plugins: [
    remix({
      unstable_ssr: false,
    }),
    vanillaExtractPlugin(),
  ],
  resolve: {
    alias: {
      '@/': path.join(__dirname, './app/')
    }
  },
})
