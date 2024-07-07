import react from "@vitejs/plugin-react-swc";
import { resolve } from "path";
import { defineConfig } from "vite";

const root = resolve(__dirname, "./web/pages/");
const outDir = resolve(__dirname, "./dist/vite");

export const hash = Math.floor(Math.random() * 90000) + 10000;

// https://vitejs.dev/config/
export default defineConfig({
  root,
  base: "/dist/vite/assets/",
  plugins: [react()],
  build: {
    minify: false,
    manifest: true,
    cssCodeSplit: false,
    outDir,
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: resolve(__dirname, "./web/pages/projects/main.tsx"),
        project: resolve(__dirname, "./web/pages/project/main.tsx"),
        editor: resolve(__dirname, "./web/pages/editor/main.tsx"),
        generation: resolve(__dirname, "./web/pages/generation/main.tsx"),
      },
      treeshake: true,
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name]` + hash + `.js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
  },
});
