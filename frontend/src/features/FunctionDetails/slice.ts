import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import fetchFunctionDetails from './fetchFunctionDetails';

interface FunctionDetailsState {
  function?: Fn;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: FunctionDetailsState = {
  errorMessage: '',
};

const slice = createSlice({
  name: 'functionDetails',
  initialState,
  reducers: {},
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

export default slice.reducer;
