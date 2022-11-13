import http, {AppError, AppResponse} from '@/utils/axios';
import {FormControl, FormLabel, FormHelperText, Input, Button, Checkbox, Textarea} from '@chakra-ui/react';
import {useQuery} from '@tanstack/react-query';
import {useState} from 'react';
import TabLayout from '../TabLayout';

const expected = {
  ok: true,
  result: {
    path: 'Sample Page-1668333114744584000',
    url: 'http://localhost:8080/Sample Page-1668333114744584000',
    title: 'Sample Page',
    description: 'Optional description',
    author_name: 'Derek Wang',
    author_url: 'https://github.com/wghglory/go-telegra',
    content: [
      {
        tag: 'p',
        children: ['Hello World'],
      },
    ],
    views: 1,
    can_edit: true,
  },
};

{
  /* http://localhost:8080/editPage?access_token=eec0b228c9fa28cc7dfd8dfbb84d47ad16a2d39e1ebf8c4f466d1acf9b443840&title=First&author_name=Derek+Wang&content=[%7B%22tag%22:%22p%22,%22children%22:[%22Hello,+world!%22]%7D]&return_content=true */
}
export default function EditPage() {
  const path = localStorage.getItem('telegra_path');
  const [api, setApi] = useState('');
  const [title, setTitle] = useState('Sample Page');
  const [authorName, setAuthorName] = useState('Derek Wang');
  const [authorUrl, setAuthorUrl] = useState('https://github.com/wghglory/go-telegra');
  const [description, setDescription] = useState('');
  const [returnContent, setReturnContent] = useState(false);
  const content = `[{"tag": "p", "children": ["You updated content!"]}]`;

  const {data, error, refetch, remove} = useQuery<AppResponse, AppError>(
    ['editPage'],
    () => {
      const params = {
        title,
        description,
        author_name: authorName,
        author_url: authorUrl,
        return_content: returnContent.toString(),
        access_token: localStorage.getItem('telegra_access_token') || '',
        content,
      };

      const paramString = new URLSearchParams(params).toString();

      setApi(encodeURI(`/editPage/${path}?${paramString}`));

      return http.get(`/editPage/${path}`, {
        params,
      });
    },
    {
      enabled: false,
      onSuccess({result}) {
        localStorage.setItem('telegra_path', result.path);
      },
    },
  );

  function submit() {
    remove();
    refetch();
  }

  return (
    <TabLayout data={data || error} api={api} expected={expected}>
      <FormControl>
        <FormLabel>Title (Required)</FormLabel>
        <Input type="text" value={title} onChange={e => setTitle(e.target.value)} />
      </FormControl>
      <FormControl>
        <FormLabel>Author Name</FormLabel>
        <Input type="text" value={authorName} onChange={e => setAuthorName(e.target.value)} />
      </FormControl>
      <FormControl>
        <FormLabel>Author Url</FormLabel>
        <Input type="text" value={authorUrl} onChange={e => setAuthorUrl(e.target.value)} />
      </FormControl>
      <FormControl>
        <FormLabel>Description</FormLabel>
        <Textarea value={description} onChange={e => setDescription(e.target.value)} />
      </FormControl>
      <FormControl>
        <Checkbox onChange={e => setReturnContent(e.target.checked)}>Return Content</Checkbox>
      </FormControl>
      <Button colorScheme="telegram" variant="solid" onClick={submit}>
        Edit Page
      </Button>
    </TabLayout>
  );
}
