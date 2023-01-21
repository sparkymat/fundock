/* eslint-disable react/no-unknown-property */
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, useParams } from 'react-router-dom';
import fetchInvocation from '../../features/InvocationsDetails/fetchInvocation';
import {
  selectInvocation,
  selectInvocationDetailsLoading,
} from '../../features/InvocationsDetails/selects';
import { AppDispatch } from '../../store';

const FunctionDetails = () => {
  const { id } = useParams();

  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    dispatch(fetchInvocation(id || ''));
  }, []);

  const inv = useSelector(selectInvocation);
  const loading = useSelector(selectInvocationDetailsLoading);

  return (
    <div className="uk-padding uk-flex uk-flex-column">
      {!loading && inv && (
        <>
          <h3>{inv.id}</h3>
          <table className="uk-table uk-table-striped uk-width-1-1">
            <tbody>
              <tr>
                <th>Function</th>
                <td>
                  <Link to={`/fn/${inv.function_name}`}>
                    {inv.function_name}
                  </Link>
                </td>
              </tr>
              <tr>
                <th>Image</th>
                <td>
                  <span uk-icon="info" />
                  {inv.image}
                </td>
              </tr>
              <tr>
                <th>Client</th>
                <td>{inv.client_name}</td>
              </tr>
            </tbody>
          </table>
          <h3>Input</h3>
          {inv.input && <pre>{inv.input}</pre>}
          {!inv && <p>No input</p>}
          <h3>Output</h3>
          {inv.output && <pre>{inv.output}</pre>}
          {!inv.output && <p>No output</p>}
        </>
      )}
    </div>
  );
};

export default FunctionDetails;
