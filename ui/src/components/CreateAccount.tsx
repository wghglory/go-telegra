import http, {AppError, AppResponse} from '@/utils/axios';
import {FormControl, FormLabel, FormHelperText, Input, Box, VStack, Button, HStack, Heading} from '@chakra-ui/react';
import {useQuery} from '@tanstack/react-query';
import {useState} from 'react';
import TabLayout from './TabLayout';

{
  /* http://localhost:8080/createAccount?short_name=Derek&author_name=Derek%20Wang&author_url=https://github.com/wghglory/go-telegra */
}
export default function CreateAccount() {
  const expected = {
    ok: true,
    result: {
      short_name: 'Derek',
      author_name: 'Derek Wang',
      author_url: 'https://github.com/wghglory/go-telegra',
      access_token: 'f7281879a3479e951c9e3aa04d9d91d7e227b15ab8dda2f34ad42ef7e08b5c69',
      auth_url: 'http://localhost:8080/auth/f7281879a3479e951c9e3aa04d9d91d7e227b15ab8dda2f34ad42ef7e08b5c69',
    },
  };
  const [shortName, setShortName] = useState('Derek');
  const [authorName, setAuthorName] = useState('Derek Wang');
  const [authorUrl, setAuthorUrl] = useState('https://github.com/wghglory/go-telegra');

  const {data, error, refetch, remove} = useQuery<AppResponse, AppError>(
    ['createAccount'],
    () =>
      http.get('/createAccount', {
        params: {
          short_name: shortName,
          author_name: authorName,
          author_url: authorUrl,
        },
      }),
    {enabled: false},
  );

  function submit() {
    remove();
    refetch();
  }

  return (
    <TabLayout data={data || error} expected={expected}>
      <FormControl>
        <FormLabel>Short Name (Required)</FormLabel>
        <Input type="text" value={shortName} onChange={e => setShortName(e.target.value)} />
        <FormHelperText>short name of the author.</FormHelperText>
      </FormControl>
      <FormControl>
        <FormLabel>Author Name</FormLabel>
        <Input type="text" value={authorName} onChange={e => setAuthorName(e.target.value)} />
        <FormHelperText>author full name.</FormHelperText>
      </FormControl>
      <FormControl>
        <FormLabel>Author URL</FormLabel>
        <Input type="text" value={authorUrl} onChange={e => setAuthorUrl(e.target.value)} />
        <FormHelperText>optional author url, i.e github link.</FormHelperText>
      </FormControl>
      <Button colorScheme="telegram" variant="solid" onClick={submit}>
        Create Account
      </Button>
    </TabLayout>
  );
}
