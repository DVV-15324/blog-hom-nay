import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';
export default defineConfig({
  plugins: [react(), tailwindcss()],
  build: {
    // 👇 Nếu build để deploy, thì đừng quên kiểm tra cấu hình server nơi deploy
    outDir: 'dist',
  },
  server: {
    proxy: {
      '/upload-image': {
        target: 'http://image.bloghomnay.com',
        changeOrigin: true,
      },
    },
  },
});