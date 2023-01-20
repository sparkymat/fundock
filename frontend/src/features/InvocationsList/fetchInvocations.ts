import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Invocation from '../../models/Invocation';

export interface FetchInvocationsRequest {
  page_size: number;
  page_number: number;
  fn: string;
}

export interface FetchInvocationsResponse {
  page_size: number;
  page_number: number;
  items: Invocation[];
}

export interface ErrorResponse {
  error: string;
}

const fetchInvocations = createAsyncThunk<
  FetchInvocationsResponse | ErrorResponse,
  FetchInvocationsRequest
>('features/fetchInvocationsList', async (request: FetchInvocationsRequest) => {
  let url = `/api/invocations?page_size=${request.page_size}&page_number=${request.page_number}`;
  if (request.fn) {
    url += `&fn=${request.fn}`;
  }
  try {
    const response = await axios.get(url);
    return response.data as FetchInvocationsResponse;
  } catch (error) {
    // eslint-disable-next-line no-console
    console.error(error);
    return { error };
  }
});

export default fetchInvocations;
