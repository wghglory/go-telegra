import {Box, Heading, VStack} from '@chakra-ui/react';

export default function ApiResult({data, expected, api}: {data: any; expected: any; api: string}) {
  return (
    <VStack align="start" spacing={8}>
      <section>
        <Heading as="h2" fontSize={24}>
          Expected Sample Result
        </Heading>
        <Box as="pre">{JSON.stringify(expected, null, 2)}</Box>
      </section>
      <section>
        <Heading as="h2" fontSize={24}>
          Your Result
        </Heading>
        {data && <Box as="pre">{JSON.stringify(data, null, 2)}</Box>}
      </section>
      <Box>
        <Heading as="h2" fontSize={24}>
          Your Network Request URL
        </Heading>
        <Box wordBreak={'break-all'}>{api}</Box>
      </Box>
    </VStack>
  );
}
