import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import fetchFunctions, {
  FetchFunctionsResponse,
  ErrorResponse,
} from './fetchFunctions';

interface FunctionsListState {
  functions: Fn[];
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: FunctionsListState = {
  functions: [],
  errorMessage: '',
};

const slice = createSlice({
  name: 'functionsList',
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder.addCase(fetchFunctions.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(fetchFunctions.fulfilled, (state, action) => {
      state.loading = false;
      state.functions = (action.payload as FetchFunctionsResponse).items;
    });
    builder.addCase(fetchFunctions.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export default slice.reducer;
