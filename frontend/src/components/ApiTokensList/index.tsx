import React, { ChangeEvent, useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import createApiToken from '../../features/ApiTokens/createApiToken';
import fetchApiTokens from '../../features/ApiTokens/fetchApiTokens';
import {
  selectApiTokens,
  selectApiTokensLoading,
  selectNewClientName,
} from '../../features/ApiTokens/selects';
import { setNewClientName } from '../../features/ApiTokens/slice';
import { AppDispatch } from '../../store';

const ApiTokensList = () => {
  const loading = useSelector(selectApiTokensLoading);
  const apiTokens = useSelector(selectApiTokens);

  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    dispatch(fetchApiTokens({ page_number: 1, page_size: 20 }));
  }, []);

  const newClientName = useSelector(selectNewClientName);
  const clientNameUpdated = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(setNewClientName(evt.target.value));
    },
    [],
  );

  const createApiTokenClicked = useCallback(() => {
    dispatch(createApiToken(newClientName));
  }, [dispatch, newClientName]);

  return (
    <div className="uk-padding">
      <div className="uk-margin-top">
        <div className="uk-flex uk-flex-row">
          <div className="uk-form-controls uk-flex-1">
            <input
              className="uk-input"
              type="text"
              placeholder="Client name (e.g. iOS app)"
              value={newClientName}
              onChange={clientNameUpdated}
            />
          </div>
          <button
            type="button"
            className="uk-button uk-button-primary uk-margin-left"
            onClick={createApiTokenClicked}
            disabled={newClientName === ''}
          >
            Create API Token
          </button>
        </div>
      </div>
      <table className="uk-table uk-table-striped">
        <thead>
          <tr>
            <th>Client Name</th>
            <th>Token</th>
            <th>Last used</th>
          </tr>
        </thead>
        <tbody>
          {apiTokens &&
            apiTokens.map(f => (
              <tr>
                <td>{f.client_name}</td>
                <td>{f.token}</td>
                <td>{f.last_used_time}</td>
              </tr>
            ))}
        </tbody>
      </table>
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

export default ApiTokensList;
