import {Box, Heading, HStack, VStack} from '@chakra-ui/react';

export default function TabLayout({data, children, expected}: any) {
  return (
    <HStack spacing={24} align="start">
      <VStack as="form" spacing={6} width="400px" align="start">
        {children}
      </VStack>
      <VStack align="start" spacing={8}>
        <section>
          <Heading as="h2" fontSize={24}>
            Expected Successful Sample Result
          </Heading>
          <Box as="pre">{JSON.stringify(expected, null, 2)}</Box>
        </section>
        <section>
          <Heading as="h2" fontSize={24}>
            Your Result
          </Heading>
          {data && <Box as="pre">{JSON.stringify(data, null, 2)}</Box>}
        </section>
      </VStack>
    </HStack>
  );
}
