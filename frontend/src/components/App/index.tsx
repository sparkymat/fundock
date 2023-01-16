import React from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import FunctionsList from '../FunctionsList';
import Home from '../Home';

const App = () => (
  <div>
    <nav className="uk-navbar-container" uk-navbar>
      <div className="uk-navbar-left">
        <Link to="/" className="uk-navbar-item uk-logo uk-margin-small-left">
          fundock
        </Link>
        <ul className="uk-navbar-nav">
          <li className="uk-parent">
            <Link to="/functions">Functions</Link>
          </li>
        </ul>
      </div>
    </nav>
    <Routes>
      <Route index element={<Home />} />
      <Route path="/functions" element={<FunctionsList />} />
    </Routes>
  </div>
);

export default App;
