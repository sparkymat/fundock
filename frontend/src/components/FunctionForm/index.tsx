import React, { ChangeEvent, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import createFunction from '../../features/FunctionForm/createFunction';
import {
  selectImage,
  selectName,
  selectSkipLogging,
  selectFormProcessing,
  selectEnvironment,
  selectSecrets,
} from '../../features/FunctionForm/selects';
import {
  addEnvironmentKey,
  addSecretsKey,
  setEnvironmentKey,
  setEnvironmentValue,
  setImage,
  setName,
  setSecretsKey,
  setSecretsValue,
  setSkipLogging,
} from '../../features/FunctionForm/slice';
import { AppDispatch } from '../../store';

const FunctionForm = () => {
  const name = useSelector(selectName);
  const image = useSelector(selectImage);
  const skipLogging = useSelector(selectSkipLogging);
  const environment = useSelector(selectEnvironment);
  const secrets = useSelector(selectSecrets);

  const navigate = useNavigate();
  const dispatch = useDispatch<AppDispatch>();

  const nameChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(setName(evt.target.value));
    },
    [dispatch],
  );

  const imageChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(setImage(evt.target.value));
    },
    [dispatch],
  );

  const skipLoggingChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(setSkipLogging(evt.target.checked));
    },
    [dispatch],
  );

  const addEnvironmentVariable = useCallback(() => {
    dispatch(addEnvironmentKey());
  }, [dispatch]);

  const environmentKeyChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      const position = parseInt(evt.target.dataset.position || '0');
      dispatch(setEnvironmentKey({ key: evt.target.value, position }));
    },
    [],
  );

  const environmentValueChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      const position = parseInt(evt.target.dataset.position || '0');
      dispatch(setEnvironmentValue({ value: evt.target.value, position }));
    },
    [],
  );

  const addSecretsVariable = useCallback(() => {
    dispatch(addSecretsKey());
  }, [dispatch]);

  const secretsKeyChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      const position = parseInt(evt.target.dataset.position || '0');
      dispatch(setSecretsKey({ key: evt.target.value, position }));
    },
    [],
  );

  const secretsValueChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      const position = parseInt(evt.target.dataset.position || '0');
      dispatch(setSecretsValue({ value: evt.target.value, position }));
    },
    [],
  );

  const submitForm = useCallback(() => {
    dispatch(
      createFunction({
        name,
        image,
        skip_logging: skipLogging,
        navigate,
        environment,
        secrets,
      }),
    );
  }, [dispatch, image, name, navigate, skipLogging, environment, secrets]);

  const formProcessing = useSelector(selectFormProcessing);

  return (
    <div className="uk-padding uk-flex uk-flex-column">
      <h3 className="uk-text-center">Add new function</h3>
      <div>
        <input type="hidden" name="csrf" value="{%s csrfToken %}" />
        <div className="uk-form-controls">
          <input
            className="uk-input"
            id="name"
            type="text"
            name="name"
            placeholder="Enter name e.g. hello"
            value={name}
            onChange={nameChanged}
            required
          />
        </div>
        <div className="uk-form-controls uk-margin-top">
          <input
            className="uk-input"
            id="image"
            type="text"
            name="image"
            placeholder="Enter image. e.g docker.io/hello-world:latest"
            value={image}
            onChange={imageChanged}
            required
          />
        </div>
        <div className="uk-margin uk-grid-small uk-child-width-auto uk-grid">
          <label htmlFor="skip_logging">
            <input
              className="uk-checkbox"
              type="checkbox"
              name="skip_logging"
              id="skip_logging"
              checked={skipLogging}
              onChange={skipLoggingChanged}
            />{' '}
            Skip logging?
          </label>
        </div>
        <h3>Environment</h3>
        <table className="uk-table uk-table-striped uk-width-1-1 uk-width-1-2@m uk-width-1-3@l">
          <thead>
            <th>Key</th>
            <th>Value</th>
          </thead>
          <tbody>
            {environment.map((kv, i) => (
              <tr key={`env-row-${i}`}>
                <td>
                  <input
                    type="text"
                    className="uk-input"
                    placeholder="Key"
                    value={kv.key}
                    data-position={i}
                    onChange={environmentKeyChanged}
                  />
                </td>
                <td>
                  <input
                    type="text"
                    className="uk-input"
                    placeholder="Value"
                    value={kv.value}
                    data-position={i}
                    onChange={environmentValueChanged}
                  />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        <button
          type="button"
          className={`uk-button uk-button-primary uk-margin-top`}
          onClick={addEnvironmentVariable}
        >
          Add environment variable
        </button>
        <h3>Secrets</h3>
        <table className="uk-table uk-table-striped uk-width-1-1 uk-width-1-2@m uk-width-1-3@l">
          <thead>
            <th>Key</th>
            <th>Value</th>
          </thead>
          <tbody>
            {secrets.map((kv, i) => (
              <tr key={`secrets-row-${i}`}>
                <td>
                  <input
                    type="text"
                    className="uk-input"
                    placeholder="Key"
                    value={kv.key}
                    data-position={i}
                    onChange={secretsKeyChanged}
                  />
                </td>
                <td>
                  <input
                    type="password"
                    className="uk-input"
                    placeholder="Value"
                    value={kv.value}
                    data-position={i}
                    onChange={secretsValueChanged}
                  />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        <button
          type="button"
          className={`uk-button uk-button-primary uk-margin-top`}
          onClick={addSecretsVariable}
        >
          Add secrets variable
        </button>
        <p></p>
        <button
          type="button"
          className={`uk-button ${
            formProcessing ? 'uk-button-default' : 'uk-button-primary'
          } uk-float-right uk-margin-top`}
          disabled={formProcessing}
          onClick={submitForm}
        >
          Add
        </button>
      </div>
    </div>
  );
};

export default FunctionForm;
