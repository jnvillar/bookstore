import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { Home } from "./components/screens/Home/Home";
import React, { useEffect, useState } from 'react';
import { Header } from "./components/header/Header";
import { Contact } from "./components/contact/Contact";

require('typeface-allura');

function App() {

  useEffect(() => {
    document.title = "La librería"
  }, []);

  const [showContact, setShowContact] = useState(false);

  return (
    <main className={"page"}>
      <Router>
        <Header showContact={setShowContact}/>
        <Contact shouldShow={showContact} handleClose={() => setShowContact(false)}/>
        <Switch>
          <Route path="/" exact>
            <Home showContact={setShowContact}/>
          </Route>
        </Switch>
      </Router>
    </main>
  );
}

export default App;