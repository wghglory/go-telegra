import {Tabs, TabList, TabPanels, Tab, TabPanel, Heading, Box} from '@chakra-ui/react';
import './App.css';
import CreateAccount from './components/CreateAccount';
import GetAccount from './components/GetAccount';

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
            <p>three!</p>
          </TabPanel>
          <TabPanel>
            <p>three!</p>
          </TabPanel>
          <TabPanel>
            <p>three!</p>
          </TabPanel>
          <TabPanel>
            <p>three!</p>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}

export default App;
