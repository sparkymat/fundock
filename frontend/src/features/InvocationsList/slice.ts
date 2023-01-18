import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Invocation from '../../models/Invocation';
import fetchInvocations, {
  ErrorResponse,
  FetchInvocationsResponse,
} from './fetchInvocations';

interface FunctionsListState {
  invocations: Invocation[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: FunctionsListState = {
  invocations: [],
  errorMessage: '',
};

const slice = createSlice({
  name: 'invocationsList',
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder.addCase(fetchInvocations.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(fetchInvocations.fulfilled, (state, action) => {
      state.loading = false;
      state.invocations = (action.payload as FetchInvocationsResponse).items;
    });
    builder.addCase(fetchInvocations.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export default slice.reducer;
