import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), '');

  return {
    plugins: [react()],
    resolve: {
      alias: {
        '~@gooddata': '/node_modules/@gooddata',
      },
    },
    server: {
      port: 8081,
      fs: {
        strict: false,
      },
      proxy: {
        '/api': {
          changeOrigin: true,
          cookieDomainRewrite: 'localhost',
          secure: false,
          target: env.VITE_HOSTNAME,
          headers: {
            host: env.VITE_HOSTNAME!,
          },
          configure: (proxy) => {
            proxy.on('proxyReq', (proxyReq) => {
              // changeOrigin: true does not work well for POST requests, so remove origin like this to be safe
              proxyReq.removeHeader('origin');
              proxyReq.setHeader('accept-encoding', 'identity');
            });
          },
        },
      },
    },
  };
});
