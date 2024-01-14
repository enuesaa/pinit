import { unstable_vitePlugin as remix } from '@remix-run/dev'
import { defineConfig } from 'vite'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'

export default defineConfig({
  plugins: [
    remix({
      unstable_ssr: false,
    }),
    vanillaExtractPlugin(),
  ],
})
