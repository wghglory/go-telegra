import React from 'react';
import * as ReactDOM from 'react-dom/client';
import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import {ReactQueryDevtools} from '@tanstack/react-query-devtools';
import './index.css';
import App from './App';

import {ChakraProvider} from '@chakra-ui/react';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // useErrorBoundary: true,
      refetchOnWindowFocus: false,
      retry: false,
      suspense: false,
      // retry(failureCount, error: any) {
      //   // retry once
      //   if (failureCount < 1) return true;
      //   else return false;
      // },
    },
  },
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <ChakraProvider>
        <App />
      </ChakraProvider>
      <ReactQueryDevtools initialIsOpen={false} position="bottom-right" />
    </QueryClientProvider>
  </React.StrictMode>,
);
