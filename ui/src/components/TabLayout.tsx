import {HStack, VStack} from '@chakra-ui/react';
import React from 'react';
import ApiResult from './ApiResult';

export default function TabLayout({
  api = '',
  data,
  expected,
  children,
  result,
}: {
  api?: string;
  data?: any;
  expected?: any;
  children: React.ReactNode;
  result?: React.ReactNode;
}) {
  return (
    <HStack spacing={24} align="start">
      <VStack as="form" spacing={6} width="400px" align="start">
        {children}
      </VStack>
      {result ? result : <ApiResult data={data} expected={expected} api={api} />}
    </HStack>
  );
}
