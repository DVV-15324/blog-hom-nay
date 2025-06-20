<<<<<<< HEAD
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
=======
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
})
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
