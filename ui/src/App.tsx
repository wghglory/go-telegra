import {Tabs, TabList, TabPanels, Tab, TabPanel, Heading, Box} from '@chakra-ui/react';
import './App.css';
import CreateAccount from './components/account/CreateAccount';
import GetAccount from './components/account/GetAccount';
import RevokeAccessToken from './components/account/RevokeAccessToken';
import CreatePage from './components/page/CreatePage';
import EditPage from './components/page/EditPage';
import GetPage from './components/page/GetPage';

function App() {
  return (
    <Box p="8">
      <Heading pb="4">Telegra Testing Client</Heading>
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
            <CreateAccount />
          </TabPanel>
          <TabPanel>
            <GetAccount />
          </TabPanel>
          <TabPanel>
            <RevokeAccessToken />
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
