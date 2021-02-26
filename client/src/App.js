import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { Home } from "./components/screens/Home/Home";
import React, { useEffect } from 'react';
import { Header } from "./components/header/Header";

function App() {

  useEffect(() => {
    document.title = "La librería"
  }, []);

  return (
    <main className={"page"}>
      <Router>
        <Header/>
        <Switch>
          <Route path="/" exact>
            <Home/>
          </Route>
        </Switch>
      </Router>
    </main>
  );
}

export default App;