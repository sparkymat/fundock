import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Invocation from '../../models/Invocation';

export interface ErrorResponse {
  error: string;
}

const fetchInvocations = createAsyncThunk<Invocation | ErrorResponse, string>(
  'features/fetchInvocation',
  async (id: string) => {
    const url = `/api/invocations/${id}`;
    try {
      const response = await axios.get(url);
      return response.data as Invocation;
    } catch (error) {
      // eslint-disable-next-line no-console
      console.error(error);
      return { error };
    }
  },
);

export default fetchInvocations;
