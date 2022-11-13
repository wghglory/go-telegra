import http, {AppError, AppResponse} from '@/utils/axios';
import {Button, Checkbox} from '@chakra-ui/react';
import {useQuery} from '@tanstack/react-query';
import {useState} from 'react';
import TabLayout from '../TabLayout';

const expected = {
  ok: true,
  result: {
    path: 'Simple Page-1668337027630672000',
    url: 'http://localhost:8080/Simple Page-1668337027630672000',
    title: 'Simple Page',
    description: 'Optional description',
    author_name: 'Derek Wang',
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
  /* http://localhost:8080/getPage?access_token=d3b25feccb89e508a9114afb82aa421fe2a9712b963b387cc5ad71e58722&fields=[%22short_name%22,%22page_count%22] */
}
export default function GetPage() {
  const path = localStorage.getItem('telegra_path');
  const [api, setApi] = useState('');
  const [returnContent, setReturnContent] = useState(false);

  const {data, error, refetch, remove} = useQuery<AppResponse, AppError>(
    ['getPage'],
    () => {
      const params = {
        return_content: returnContent.toString(),
      };

      const paramString = new URLSearchParams(params).toString();
      setApi(encodeURI(`/getPage/${path}?${paramString}`));

      return http.get(`/getPage/${path}`, {
        params,
      });
    },
    {enabled: false},
  );

  function getPage() {
    remove();
    refetch();
  }

  return (
    <TabLayout api={api} data={data || error} expected={expected}>
      <Checkbox onChange={e => setReturnContent(e.target.checked)}>Return Content</Checkbox>
      <Button colorScheme="telegram" variant="solid" onClick={getPage}>
        Get Page
      </Button>
    </TabLayout>
  );
}
