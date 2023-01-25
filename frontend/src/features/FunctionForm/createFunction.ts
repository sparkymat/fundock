import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';

interface KeyValuePair {
  key: string;
  value: string;
}

interface CreateFunctionRequest {
  name: string;
  image: string;
  skip_logging: boolean;
  environment: KeyValuePair[];
  secrets: KeyValuePair[];
  navigate(_: string): void;
}

const createFunction = createAsyncThunk<
  Fn | ErrorResponse,
  CreateFunctionRequest
>('features/fetchFunctionDetails', async (request: CreateFunctionRequest) => {
  try {
    const environment: { [key: string]: string } = {};
    request.environment.forEach(kv => {
      environment[kv.key] = kv.value;
    });

    const secrets: { [key: string]: string } = {};
    request.secrets.forEach(kv => {
      secrets[kv.key] = kv.value;
    });

    console.log(environment);
    console.log(secrets);

    const response = await axios.post('/api/functions', {
      name: request.name,
      image: request.image,
      skip_logging: request.skip_logging,
      environment,
      secrets,
    });
    request.navigate(`/fn/${request.name}`);
    return response.data as Fn;
  } catch (error) {
    console.error(error);
    return { error };
  }
});

export default createFunction;
