// go-ecommerce-frontend/vite.config.js
import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  root: './', // Serve from the project root
  build: {
    rollupOptions: {
      input: {
        main: resolve('go-ecommerce-frontend', 'index.html'),
        login: resolve('go-ecommerce-frontend', 'src/pages/login.html'),
        register: resolve('go-ecommerce-frontend', 'src/pages/register.html'),
        productDetail: resolve('go-ecommerce-frontend', 'src/pages/product-detail.html'),
        cart: resolve('go-ecommerce-frontend', 'src/pages/cart.html'),
        checkout: resolve('go-ecommerce-frontend', 'src/pages/checkout.html'),
        profile: resolve('go-ecommerce-frontend', 'src/pages/profile.html'),
        adminProducts: resolve('go-ecommerce-frontend', 'src/pages/admin/products.html'),
        adminAddProduct: resolve('go-ecommerce-frontend', 'src/pages/admin/add-product.html'),
        adminEditProduct: resolve('go-ecommerce-frontend', 'src/pages/admin/edit-product.html'),
      },
    },
    outDir: 'dist', // Output directory for build
  },
  server: {
    port: 3000, // Port for the Vite dev server
    open: true,   // Open browser on start
  },
});