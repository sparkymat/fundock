import React, { ChangeEvent, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import createFunction from '../../features/FunctionForm/createFunction';
import {
  selectImage,
  selectName,
  selectSkipLogging,
  selectFormProcessing,
} from '../../features/FunctionForm/selects';
import {
  setImage,
  setName,
  setSkipLogging,
} from '../../features/FunctionForm/slice';
import { AppDispatch } from '../../store';

const FunctionForm = () => {
  const name = useSelector(selectName);
  const image = useSelector(selectImage);
  const skipLogging = useSelector(selectSkipLogging);

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

  const submitForm = useCallback(() => {
    dispatch(
      createFunction({
        name,
        image,
        skip_logging: skipLogging,
        navigate,
      }),
    );
  }, [dispatch, image, name, navigate, skipLogging]);

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
