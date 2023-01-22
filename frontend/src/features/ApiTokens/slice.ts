import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import ApiToken from '../../models/ApiToken';
import Fn from '../../models/Fn';
import Invocation from '../../models/Invocation';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import createApiToken from './createApiToken';
import fetchApiTokens, { FetchApiTokensResponse } from './fetchApiTokens';

interface ApiTokensState {
  apiTokens: ApiToken[];
  errorMessage: string;
  newClientName: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: ApiTokensState = {
  apiTokens: [],
  errorMessage: '',
  newClientName: '',
};

const slice = createSlice({
  name: 'apiTokens',
  initialState,
  reducers: {
    setNewClientName: (state, action: PayloadAction<string>) => {
      state.newClientName = action.payload;
    },
  },
  extraReducers: builder => {
    builder.addCase(fetchApiTokens.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(fetchApiTokens.fulfilled, (state, action) => {
      state.loading = false;
      state.apiTokens = (action.payload as FetchApiTokensResponse).items;
    });
    builder.addCase(fetchApiTokens.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
    builder.addCase(createApiToken.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(createApiToken.fulfilled, (state, action) => {
      state.loading = false;
    });
    builder.addCase(createApiToken.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export const { setNewClientName } = slice.actions;

export default slice.reducer;
