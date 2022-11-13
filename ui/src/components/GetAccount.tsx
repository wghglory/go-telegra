import http, {AppError, AppResponse} from '@/utils/axios';
import {Box, Button, Stack, Checkbox, Divider} from '@chakra-ui/react';
import {useQuery} from '@tanstack/react-query';
import {useMemo, useState} from 'react';
import ApiResult from './ApiResult';
import TabLayout from './TabLayout';

const expected1 = {
  ok: true,
  result: {
    short_name: 'Derek',
    author_name: 'Derek Wang',
    author_url: 'https://github.com/wghglory/go-telegra',
  },
};
const expected2 = {
  ok: true,
  result: {
    short_name: 'Derek',
    author_name: 'Derek Wang',
    author_url: 'https://github.com/wghglory/go-telegra',
    auth_url: 'http://localhost:8080/auth/f7281879a3479e951c9e3aa04d9d91d7e227b15ab8dda2f34ad42ef7e08b5c69',
    page_count: 1,
  },
};
const expected3 = {
  ok: false,
  error: 'wrong access token! cannot get account with this token',
};

{
  /* http://localhost:8080/getAccountInfo?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&fields=[%22short_name%22,%22page_count%22] */
}
export default function GetAccount() {
  const [api, setApi] = useState('');
  const [step, setStep] = useState(1);
  const [fields, setFields] = useState<string[]>([]);

  const fieldString = useMemo(() => {
    if (fields.length > 0) {
      return `[${fields.map(f => `"${f}"`).join(',')}]`;
    } else {
      return '';
    }
  }, [fields]);

  function changeField(value: string) {
    if (fields.includes(value)) {
      setFields(fields.filter(v => v !== value));
    } else {
      setFields([...fields, value]);
    }
  }

  const {
    data: data1,
    error: error1,
    refetch: refetch1,
    remove: remove1,
  } = useQuery<AppResponse, AppError>(
    ['getAccountInfoDefault'],
    () => {
      const params = {
        access_token: localStorage.getItem('telegra_access_token') || '',
      };

      const paramString = new URLSearchParams(params).toString();
      setApi(`/getAccountInfo/${paramString}`);

      return http.get('/getAccountInfo', {
        params,
      });
    },
    {enabled: false},
  );

  const {
    data: data2,
    error: error2,
    refetch: refetch2,
    remove: remove2,
  } = useQuery<AppResponse, AppError>(
    ['getAccountInfo'],
    () => {
      const params = {
        access_token: localStorage.getItem('telegra_access_token') || '',
        fields: fieldString,
      };

      const paramString = new URLSearchParams(params).toString();
      setApi(`/getAccountInfo/${paramString}`);

      return http.get('/getAccountInfo', {
        params,
      });
    },
    {enabled: false},
  );

  const {
    data: data3,
    error: error3,
    refetch: refetch3,
    remove: remove3,
  } = useQuery<AppResponse, AppError>(
    ['getAccountInfoWithError'],
    () => {
      const params = {
        access_token: 'a_wrong_token',
        fields: fieldString,
      };

      const paramString = new URLSearchParams(params).toString();
      setApi(`/getAccountInfo/${paramString}`);

      return http.get('/getAccountInfo', {
        params,
      });
    },
    {enabled: false},
  );

  function getAccountWithDefaultFields() {
    setStep(1);
    remove1();
    refetch1();
  }
  function getAccount() {
    setStep(2);
    remove2();
    refetch2();
  }
  function getAccountWithWrongToken() {
    setStep(3);
    remove3();
    refetch3();
  }

  return (
    <TabLayout
      result={
        step === 1 ? (
          <ApiResult api={api} data={data1 || error1} expected={expected1} />
        ) : step === 2 ? (
          <ApiResult api={api} data={data2 || error2} expected={expected2} />
        ) : (
          <ApiResult api={api} data={data3 || error3} expected={expected3} />
        )
      }
    >
      <Box>
        Default fields: <strong>short_name, author_name, author_url</strong>
      </Box>
      <Button colorScheme="telegram" variant="solid" onClick={getAccountWithDefaultFields}>
        Get Account with default fields
      </Button>
      <Divider />

      <Stack spacing={2} direction="column">
        <Checkbox value="short_name" onChange={e => changeField(e.target.value)}>
          Short Name
        </Checkbox>
        <Checkbox value="author_name" onChange={e => changeField(e.target.value)}>
          Author Name
        </Checkbox>
        <Checkbox value="author_url" onChange={e => changeField(e.target.value)}>
          Author Url
        </Checkbox>
        <Checkbox value="auth_url" onChange={e => changeField(e.target.value)}>
          Auth Url
        </Checkbox>
        <Checkbox value="page_count" onChange={e => changeField(e.target.value)}>
          Page Count
        </Checkbox>
      </Stack>
      <Button colorScheme="telegram" variant="solid" onClick={getAccount}>
        Get Account with checked fields
      </Button>
      <Divider />

      <Button colorScheme="red" variant="solid" onClick={getAccountWithWrongToken}>
        Get Account with wrong access token
      </Button>
    </TabLayout>
  );
}
