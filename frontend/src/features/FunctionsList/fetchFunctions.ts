import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Fn from '../../models/Fn';

export interface FetchFunctionsRequest {
  page_size: number;
  page_number: number;
}

export interface FetchFunctionsResponse {
  page_size: number;
  page_number: number;
  items: Fn[];
}

export interface ErrorResponse {
  error: string;
}

const fetchFunctions = createAsyncThunk<
  FetchFunctionsResponse | ErrorResponse,
  FetchFunctionsRequest
>('features/fetchFunctions', async (request: FetchFunctionsRequest) => {
  try {
    const response = await axios.get(
      `/api/functions?page_size=${request.page_size}&page_number=${request.page_number}`,
    );
    return response.data as FetchFunctionsResponse;
  } catch (error) {
    console.error(error);
    return { error: error };
  }
});

export default fetchFunctions;
