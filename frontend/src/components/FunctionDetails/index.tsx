import { findNonSerializableValue } from '@reduxjs/toolkit';
import React from 'react';
import { useParams } from 'react-router-dom';

const FunctionDetails = () => {
  const { name } = useParams();
  
  return (
    <div className="uk-padding uk-flex uk-flex-column">
    <h1>{ fn.name }</h1>
    <p className="uk-margin-small">
      <span uk-icon="info"></span>
      { fn.image }
    </p>
    { fn.skip_logging &&      <p>Input and output is not logged</p> }
    { !fn.skip_logging &&      <p>Input and output is logged</p> }
    <div>
      <button className="uk-button uk-button-default">Disable</button>
    </div>
    <div className="uk-margin-top uk-width-1-1 uk-width-1-2@m">
      <form action="/exec/{%s fn.Name %}" method="POST">
        <input type="hidden" name="csrf" value="{%s csrfToken %}" />
        <textarea className="uk-width-1-1 uk-textarea" rows="8" name="input">{}</textarea>
        <input type="submit" className="uk-button uk-button-primary uk-margin-small-top" value="Run" />
      </form>
    </div>
    <h3>Invocations</h3>
    <table className="uk-table uk-table-striped">
      <thead>
        <th>ID</th>
        <th>Client</th>
        <th>Status</th>
        <th>Timestamp</th>
        <th>Duration</th>
      </thead>
      <tbody>
          <tr>
            <td><a href="/invocations/{%s in.ID %}">{ in.ID }</a></td>
            <td>{ in.ClientName %}</td>
            <td>{ in.Status %}</td>
            <td>{ in.Timestamp %}</td>
            <td>{ in.Duration %}</td>
          </tr>
      </tbody>
    </table>
  </div>

  );
};

export default FunctionDetails;
