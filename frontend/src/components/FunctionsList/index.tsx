import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { selectFunctionsListLoading } from '../../features/FunctionsList/selects';
import { fetchFunctions } from '../../features/FunctionsList/slice';

const FunctionsList = () => {
  const loading = useSelector(selectFunctionsListLoading);

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchFunctions());
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
          <tr>
            <td>
              <a href="/fn/{%s fn.Name %}">fn.Name</a>
            </td>
            <td>image</td>
            <td>No</td>
            <td>fn.CreatedTimestamp</td>
          </tr>
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
