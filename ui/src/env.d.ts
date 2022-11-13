// https://stackoverflow.com/questions/66039933/typescript-types-for-import-meta-env

interface ImportMetaEnv {
  readonly VITE_WEBSITE_NAME: string;
  readonly VITE_BASE_HREF: string;
  readonly VITE_APP_HOST: string;
  readonly VITE_APP_PREFIX: string;
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
