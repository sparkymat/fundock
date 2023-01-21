import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';

interface CreateFunctionRequest {
  name: string;
  image: string;
  skip_logging: boolean;
  navigate(_: string): void;
}

const createFunction = createAsyncThunk<
  Fn | ErrorResponse,
  CreateFunctionRequest
>('features/fetchFunctionDetails', async (request: CreateFunctionRequest) => {
  try {
    const response = await axios.post('/api/functions', {
      name: request.name,
      image: request.image,
      skip_logging: request.skip_logging ? 'true' : 'false',
    });
    request.navigate(`/fn/${request.name}`);
    return response.data as Fn;
  } catch (error) {
    console.error(error);
    return { error };
  }
});

export default createFunction;
