import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import fetchFunctionDetails from './fetchFunctionDetails';

interface FunctionDetailsState {
  function?: Fn;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  requestBody: string;
}

const initialState: FunctionDetailsState = {
  errorMessage: '',
  requestBody: '{}',
};

const slice = createSlice({
  name: 'functionDetails',
  initialState,
  reducers: {
    setRequestBody: (state, action: PayloadAction<string>) => {
      state.requestBody = action.payload;
    },
  },
  extraReducers: builder => {
    builder.addCase(fetchFunctionDetails.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(fetchFunctionDetails.fulfilled, (state, action) => {
      state.loading = false;
      state.function = action.payload as Fn;
    });
    builder.addCase(fetchFunctionDetails.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export const { setRequestBody } = slice.actions;

export default slice.reducer;
