import { findNonSerializableValue } from '@reduxjs/toolkit';
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, useParams } from 'react-router-dom';
import fetchFunctionDetails from '../../features/FunctionDetails/fetchFunctionDetails';
import {
  selectFunction,
  selectFunctionDetailsLoading,
} from '../../features/FunctionDetails/selects';
import fetchInvocations from '../../features/InvocationsList/fetchInvocations';
import {
  selectInvocations,
  selectInvocationsListLoading,
} from '../../features/InvocationsList/selects';
import Invocation from '../../models/Invocation';
import { AppDispatch } from '../../store';

const FunctionDetails = () => {
  const { name } = useParams();

  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    dispatch(fetchFunctionDetails(name || ''));
    dispatch(
      fetchInvocations({
        fn: name || '',
        page_number: 1,
        page_size: 20,
      }),
    );
  }, []);

  const fn = useSelector(selectFunction);
  const invocations = useSelector(selectInvocations);

  const loading = useSelector(selectFunctionDetailsLoading);
  const invocationsLoading = useSelector(selectInvocationsListLoading);

  const csrfToken = '';

  return (
    <div className="uk-padding uk-flex uk-flex-column">
      {fn && (
        <>
          <h1>{fn.name}</h1>
          <p className="uk-margin-small">
            <span uk-icon="info" />
            {fn.image}
          </p>
          {fn.skip_logging && <p>Input and output is not logged</p>}
          {!fn.skip_logging && <p>Input and output is logged</p>}
          <div>
            <button className="uk-button uk-button-default">Disable</button>
          </div>
          <div className="uk-margin-top uk-width-1-1 uk-width-1-2@m">
            <input type="hidden" name="csrf" value={csrfToken} />
            <textarea
              className="uk-width-1-1 uk-textarea"
              rows={8}
              name="input"
            >
              {}
            </textarea>
            <input
              type="submit"
              className="uk-button uk-button-primary uk-margin-small-top"
              value="Run"
            />
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
              {invocations &&
                invocations.map((inv: Invocation) => (
                  <tr>
                    <td>
                      <Link to={`/invocations/${inv.id}`}>{inv.id}</Link>
                    </td>
                    <td>{inv.client_name}</td>
                    <td>{inv.status}</td>
                    <td>{inv.started_time}</td>
                    <td>{inv.ended_time}</td>
                  </tr>
                ))}
            </tbody>
          </table>
          {invocationsLoading && (
            <>
              <div className="uk-overlay-default uk-position-cover" />
              <div
                className="uk-overlay uk-position-center uk-dark"
                uk-spinner="ratio: 3"
              />
            </>
          )}
        </>
      )}
      {loading && (
        <>
          <div className="uk-overlay-default uk-position-cover"></div>
          <div
            className="uk-overlay uk-position-center uk-dark"
            uk-spinner="ratio: 3"
          ></div>
        </>
      )}
    </div>
  );
};

export default FunctionDetails;
