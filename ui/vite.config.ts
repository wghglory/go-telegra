import {defineConfig} from 'vite';
import {resolve} from 'path'; // Need to install @types/node"
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  base: './',
  server: {
    port: 4000,
    open: true,
    cors: true,
    proxy: {
      // common usage
      '^/(createAccount|getAccountInfo|revokeAccessToken|getPage|createPage|editPage)': {
        target: 'http://localhost:8080', // mock
        changeOrigin: true,
        secure: false, // [vite] http proxy error: Error: self signed certificate in certificate chain
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
      // api/sessions POST, api/session, api/task GET
      '^/(api|cloudapi)/.*': {
        target: 'http://localhost:3000', // mock
        // target: 'https://alp-vcd103.eng.vmware.com',
        // target: 'https://alp-vcd104.eng.vmware.com',
        changeOrigin: true,
        secure: false, // [vite] http proxy error: Error: self signed certificate in certificate chain
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
      // regex
      '^/fallback/.*': {
        target: 'http://jsonplaceholder.typicode.com',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/fallback/, ''),
      },
    },
  },
});
