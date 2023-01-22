import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import ApiToken from '../../models/ApiToken';

export interface FetchApiTokensRequest {
  page_size: number;
  page_number: number;
}

export interface FetchApiTokensResponse {
  page_size: number;
  page_number: number;
  items: ApiToken[];
}

export interface ErrorResponse {
  error: string;
}

const fetchApiTokens = createAsyncThunk<
  FetchApiTokensResponse | ErrorResponse,
  FetchApiTokensRequest
>('features/fetchFunctions', async (request: FetchApiTokensRequest) => {
  try {
    const response = await axios.get(
      `/api/api_tokens?page_size=${request.page_size}&page_number=${request.page_number}`,
    );
    return response.data as FetchApiTokensResponse;
  } catch (error) {
    console.error(error);
    return { error: error };
  }
});

export default fetchApiTokens;
