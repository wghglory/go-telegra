import http, {AppError, AppResponse} from '@/utils/axios';
import {Button} from '@chakra-ui/react';
import {useQuery} from '@tanstack/react-query';
import {useEffect, useState} from 'react';
import TabLayout from '../TabLayout';

const expected = {
  ok: true,
  result: {
    access_token: 'f7281879a3479e951c9e3aa04d9d91d7e227b15ab8dda2f34ad42ef7e08b5c69',
    auth_url: 'http://localhost:8080/auth/f7281879a3479e951c9e3aa04d9d91d7e227b15ab8dda2f34ad42ef7e08b5c69',
  },
};

{
  /* http://localhost:8080/revokeAccessToken?short_name=Derek&author_name=Derek%20Wang&author_url=https://github.com/wghglory/go-telegra */
}
export default function revokeAccessToken({
  setGlobalToken,
}: {
  setGlobalToken: React.Dispatch<React.SetStateAction<string>>;
}) {
  const [api, setApi] = useState('');
  const [token, setToken] = useState('');

  useEffect(() => {
    if (token) {
      setTimeout(() => {
        remove();
        refetch();
      }, 50);
    }
  }, [token]);

  const {data, error, refetch, remove} = useQuery<AppResponse, AppError>(
    ['revokeAccessToken', token],
    () => {
      const params = {
        access_token: token,
      };

      const paramString = new URLSearchParams(params).toString();
      setApi(`/revokeAccessToken/${paramString}`);

      return http.get('/revokeAccessToken', {
        params,
      });
    },
    {
      enabled: false,
      onSuccess({result}) {
        localStorage.setItem('telegra_access_token', result.access_token);
        setGlobalToken(result.access_token);
      },
    },
  );

  return (
    <TabLayout data={data || error} api={api} expected={expected}>
      <Button
        colorScheme="telegram"
        variant="solid"
        onClick={() => setToken(localStorage.getItem('telegra_access_token') || '')}
      >
        Revoke Access Token
      </Button>
      <Button colorScheme="red" variant="solid" onClick={() => setToken('a_wrong_token')}>
        Revoke Access Token with wrong token
      </Button>
    </TabLayout>
  );
}
