/* eslint-disable react/no-unknown-property */
import React from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import ApiTokensList from '../ApiTokensList';
import FunctionDetails from '../FunctionDetails';
import FunctionForm from '../FunctionForm';
import FunctionsList from '../FunctionsList';
import Home from '../Home';
import InvocationDetails from '../InvocationDetails';

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
          <li className="uk-parent">
            <Link to="/api_tokens">API Tokens</Link>
          </li>
        </ul>
      </div>
    </nav>
    <Routes>
      <Route index element={<Home />} />
      <Route path="/functions" element={<FunctionsList />} />
      <Route path="/fn/:name" element={<FunctionDetails />} />
      <Route path="/invocations/:id" element={<InvocationDetails />} />
      <Route path="/functions/new" element={<FunctionForm />} />
      <Route path="/api_tokens" element={<ApiTokensList />} />
    </Routes>
  </div>
);

export default App;
