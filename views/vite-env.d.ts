/// <reference types="vite/client" />

// To prevent accidentally leaking env variables to the client, only variables prefixed with VITE_ are exposed to your Vite-processed code.
// ref: https://vitejs.dev/guide/env-and-mode
interface ImportMetaEnv {
  readonly VITE_HOSTNAME: string;
  readonly VITE_WORKSPACEID: string;
  readonly VITE_TOKEN: string;
  readonly VITE_DASHBOARDID: string;
  readonly VITE_JOUNEY_METRIC_VISUALIZATION: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
