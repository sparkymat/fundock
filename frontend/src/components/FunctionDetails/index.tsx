/* eslint-disable react/no-unknown-property */
import React, { ChangeEvent, useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link, useParams } from 'react-router-dom';
import AceEditor from 'react-ace';
import ReactModal from 'react-modal';

import fetchFunctionDetails from '../../features/FunctionDetails/fetchFunctionDetails';
import {
  selectFunction,
  selectFunctionDetailsLoading,
  selectInvocation,
  selectRequestBody,
  selectShowInvocationOutput,
} from '../../features/FunctionDetails/selects';
import {
  dismissInvocationOutput,
  setRequestBody,
} from '../../features/FunctionDetails/slice';
import fetchInvocations from '../../features/InvocationsList/fetchInvocations';
import {
  selectInvocations,
  selectInvocationsListLoading,
} from '../../features/InvocationsList/selects';
import Invocation from '../../models/Invocation';
import { AppDispatch } from '../../store';
import runFunction from '../../features/FunctionDetails/runFunction';
import { freezeDraftable } from '@reduxjs/toolkit/dist/utils';

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
  const requestBody = useSelector(selectRequestBody);

  const loading = useSelector(selectFunctionDetailsLoading);
  const invocationsLoading = useSelector(selectInvocationsListLoading);

  const invocation = useSelector(selectInvocation);
  const showInvocationOutput = useSelector(selectShowInvocationOutput);

  const requestBodyUpdated = useCallback(
    (value: string) => {
      dispatch(setRequestBody(value));
    },
    [dispatch],
  );

  const functionExecuted = useCallback(() => {
    dispatch(runFunction({ fn: fn?.name || '', requestBody }));
  }, [dispatch, fn, requestBody]);

  const invocationModalDismissed = useCallback(() => {
    dispatch(dismissInvocationOutput());
  }, []);

  const customStyles = {
    content: {
      minWidth: '50%',
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      padding: 0,
      marginRight: '-50%',
      transform: 'translate(-50%, -50%)',
    },
  };

  return (
    <div className="uk-padding uk-flex uk-flex-column">
      {fn && (
        <>
          <h1>{fn.name}</h1>
          <p className="uk-margin-small-top uk-margin-small-bottom">
            <span uk-icon="info" className="uk-margin-small-right" />
            {fn.image}
          </p>
          <h3 className="uk-margin-small-top uk-margin-small-bottom">
            Environment
          </h3>
          <table className="uk-table uk-table-striped uk-width-1-1 uk-width-1-2@m uk-width-1-3@l">
            <thead>
              <th>Key</th>
              <th>Value</th>
            </thead>
            <tbody>
              {Object.keys(fn.environment).map(k => (
                <tr>
                  <td>{k}</td>
                  <td>{fn.environment[k]}</td>
                </tr>
              ))}
            </tbody>
          </table>
          <h3 className="uk-margin-small-top uk-margin-small-bottom">
            Secrets
          </h3>
          <table className="uk-table uk-table-striped uk-width-1-1 uk-width-1-2@m uk-width-1-3@l">
            <thead>
              <th>Key</th>
              <th>Value</th>
            </thead>
            <tbody>
              {fn.secrets.map(k => (
                <tr>
                  <td>{k}</td>
                  <td>****</td>
                </tr>
              ))}
            </tbody>
          </table>
          {fn.skip_logging && <p>Input and output is not logged</p>}
          {!fn.skip_logging && <p>Input and output is logged</p>}
          <div>
            <button type="button" className="uk-button uk-button-default">
              Disable
            </button>
          </div>
          <div className="uk-margin-top uk-width-1-1 uk-width-1-2@m">
            <AceEditor
              mode="json"
              theme="solarized_dark"
              value={requestBody}
              onChange={requestBodyUpdated}
              minLines={16}
              maxLines={16}
              width="100%"
              fontSize="1.1rem"
              showGutter={false}
              editorProps={{ $blockScrolling: true }}
            />
            <button
              type="button"
              className="uk-button uk-button-primary uk-margin-small-top"
              onClick={functionExecuted}
            >
              Run
            </button>
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
                    <td>n/a</td>
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
          <div className="uk-overlay-default uk-position-cover" />
          <div
            className="uk-overlay uk-position-center uk-dark"
            uk-spinner="ratio: 3"
          />
        </>
      )}
      {invocation && (
        <ReactModal isOpen={showInvocationOutput} style={customStyles}>
          <div className="uk-card uk-card-default uk-card-body uk-width-1-1">
            <h3 className="uk-card-title">{invocation.id}</h3>
            <div className="uk-width-1-1">
              <AceEditor
                mode="json"
                theme="solarized_dark"
                value={invocation.output}
                width="100%"
                fontSize="1.1rem"
                readOnly
                showGutter={false}
                editorProps={{ $blockScrolling: true }}
              />
            </div>
            <button
              className="uk-button uk-button-primary uk-margin-top uk-align-right"
              onClick={invocationModalDismissed}
            >
              Close Modal
            </button>
          </div>
        </ReactModal>
      )}
    </div>
  );
};

export default FunctionDetails;
