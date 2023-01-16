import { findNonSerializableValue } from '@reduxjs/toolkit';
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import fetchFunctions from '../../features/FunctionsList/fetchFunctions';
import {
  selectFunctions,
  selectFunctionsListLoading,
} from '../../features/FunctionsList/selects';
import { AppDispatch } from '../../store';

const FunctionsList = () => {
  const loading = useSelector(selectFunctionsListLoading);
  const functions = useSelector(selectFunctions);

  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    dispatch(fetchFunctions({ page_number: 1, page_size: 20 }));
  }, []);

  return (
    <div className="uk-padding">
      <Link
        to="/functions/new"
        className="uk-button uk-button-primary uk-float-right"
      >
        New function
      </Link>
      <table className="uk-table uk-table-striped">
        <thead>
          <tr>
            <th>Name</th>
            <th>Image</th>
            <th>logged?</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {functions.map(f => (
            <tr>
              <td>
                <Link to={`/fn/${f.name}`}>{f.name}</Link>
              </td>
              <td>{f.image}</td>
              <td>{f.skip_logging ? 'No' : 'Yes'}</td>
              <td>{f.created_time}</td>
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

export default FunctionsList;
