import react from "@vitejs/plugin-react-swc";
import { resolve } from 'path'
import { defineConfig } from 'vite'

const root = resolve(__dirname, './web/pages/')
const outDir = resolve(__dirname, './dist/vite')

// https://vitejs.dev/config/
export default defineConfig({
  root,
  plugins: [react()],
  build: {
    minify: false,
    manifest: true,
    outDir,
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: resolve(__dirname, './web/pages/projects/main.tsx'),
      },
      treeshake: true,
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/index-chunk.js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    }
  }
})
