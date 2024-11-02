import { defineConfig } from 'vite'
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
  ],
  resolve: {
    alias: {
      '@/': path.join(__dirname, './src/')
    }
  },
})
