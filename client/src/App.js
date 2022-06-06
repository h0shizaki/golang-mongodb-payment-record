import { Route,  Routes } from "react-router-dom";
import Home from './pages/home'
import Dashboard from './pages/dashboard'

import Nav from "./components/nav";
import { Fragment } from "react";

function App() {
  return (
    <Fragment>
      <Nav />
      <Routes>
        <Route path="/" element={<Home />}></Route>
        <Route path="/dashboard" element={<Dashboard />}></Route>
      </Routes>
    </Fragment>
  );
}

export default App;
