import { createSlice } from '@reduxjs/toolkit';
import Invocation from '../../models/Invocation';
import fetchInvocation, { ErrorResponse } from './fetchInvocation';

interface InvocationDetailsState {
  invocation?: Invocation;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: InvocationDetailsState = {
  errorMessage: '',
};

const slice = createSlice({
  name: 'invocationDetails',
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder.addCase(fetchInvocation.pending, state => {
      state.loading = true;
    });
    builder.addCase(fetchInvocation.fulfilled, (state, action) => {
      state.loading = false;
      state.invocation = action.payload as Invocation;
    });
    builder.addCase(fetchInvocation.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export default slice.reducer;
