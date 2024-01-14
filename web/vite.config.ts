import { defineConfig } from 'vite'
import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin'
import react from '@vitejs/plugin-react'
import path from 'node:path'

export default defineConfig({
  root: './',
  esbuild: {
    target: 'esnext',
    format: 'esm',
  },
  build: {
    outDir: './dist',
  },
  plugins: [
    react(),
    vanillaExtractPlugin(),
  ],
  resolve: {
    alias: {
      '@/': path.join(__dirname, './src/')
    }
  },
})
