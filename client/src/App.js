import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { Home } from "./components/screens/Home/Home";
import React, { useEffect } from 'react';

function App() {

  useEffect(() => {
    document.title = "La librería"
  }, []);

  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          <Home/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;