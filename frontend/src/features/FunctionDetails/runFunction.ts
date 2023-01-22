import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Invocation from '../../models/Invocation';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';

interface RunFunctionRequest {
  fn: string;
  requestBody: string;
}

const runFunction = createAsyncThunk<
  Invocation | ErrorResponse,
  RunFunctionRequest
>('features/runFunction', async (request: RunFunctionRequest) => {
  try {
    const response = await axios.post(
      `/api/fn/${request.fn}/exec`,
      request.requestBody,
    );
    return response.data as Invocation;
  } catch (error) {
    console.error(error);
    return { error };
  }
});

export default runFunction;
