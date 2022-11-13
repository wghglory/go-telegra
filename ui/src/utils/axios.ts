import axios, {AxiosError} from 'axios';

const http = axios.create({
  // https://github.com/vitejs/vite/issues/1149 Jest testing issues.
  baseURL: `${import.meta.env.VITE_APP_HOST}/${import.meta.env.VITE_APP_PREFIX}`,
  // withCredentials: true,
  // timeout: 10000,
});

function responseErrorHandler(error: Error | AxiosError) {
  if (axios.isAxiosError(error)) {
    // Access to config, request, and response
    // timeout also here
    const {response} = error as AxiosError;
    return Promise.reject(response?.data as AppError);
  } else {
    // native error
    console.error(error, 'native error');
  }
}

http.interceptors.response.use(res => res.data, responseErrorHandler);

export interface AppResponse {
  ok: boolean;
  result: any;
}

export interface AppError {
  ok: boolean;
  error: any;
}

export default http;
