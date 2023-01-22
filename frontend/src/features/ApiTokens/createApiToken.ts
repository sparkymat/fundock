import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import ApiToken from '../../models/ApiToken';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';

const createApiToken = createAsyncThunk<ApiToken | ErrorResponse, string>(
  'features/createApiToken',
  async (client_name: string) => {
    try {
      const response = await axios.post('/api/api_tokens', {
        client_name,
      });
      window.location.reload();
      return response.data as ApiToken;
    } catch (error) {
      console.error(error);
      return { error };
    }
  },
);

export default createApiToken;
