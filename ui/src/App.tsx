import {Tabs, TabList, TabPanels, Tab, TabPanel, Heading, Box, HStack, Badge} from '@chakra-ui/react';
import {useState} from 'react';
import './App.css';
import CreateAccount from './components/account/CreateAccount';
import GetAccount from './components/account/GetAccount';
import RevokeAccessToken from './components/account/RevokeAccessToken';
import CreatePage from './components/page/CreatePage';
import EditPage from './components/page/EditPage';
import GetPage from './components/page/GetPage';

function App() {
  const [shortName, setShortName] = useState('Derek');
  const [token, setToken] = useState('');

  return (
    <Box p="8">
      <HStack justify={'space-between'}>
        <Heading pb="4">Telegra Testing Client</Heading>
        <Box fontWeight={600} fontSize="18">
          {shortName}
          <Badge colorScheme="green" ml="2" fontSize="0.8em">
            {token}
          </Badge>
        </Box>
      </HStack>
      <Tabs>
        <TabList>
          <Tab>Create Account</Tab>
          <Tab>Get AccountInfo</Tab>
          <Tab>Revoke AccessToken</Tab>
          <Tab>Create Page</Tab>
          <Tab>Get Page</Tab>
          <Tab>Edit Page</Tab>
        </TabList>

        <TabPanels>
          <TabPanel>
            <CreateAccount shortName={shortName} setShortName={setShortName} setToken={setToken} />
          </TabPanel>
          <TabPanel>
            <GetAccount token={token} />
          </TabPanel>
          <TabPanel>
            <RevokeAccessToken setGlobalToken={setToken} />
          </TabPanel>
          <TabPanel>
            <CreatePage />
          </TabPanel>
          <TabPanel>
            <GetPage />
          </TabPanel>
          <TabPanel>
            <EditPage />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}

export default App;
