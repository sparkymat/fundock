import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import Invocation from '../../models/Invocation';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import fetchFunctionDetails from './fetchFunctionDetails';
import runFunction from './runFunction';

interface FunctionDetailsState {
  function?: Fn;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
  executing?: boolean;
  requestBody: string;
  invocation?: Invocation;
  showInvocationOutput?: boolean;
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
    dismissInvocationOutput: state => {
      state.showInvocationOutput = false;
      window.location.reload();
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
    builder.addCase(runFunction.pending, (state, action) => {
      state.loading = true;
    });
    builder.addCase(runFunction.fulfilled, (state, action) => {
      state.loading = false;
      state.invocation = action.payload as Invocation;
      state.showInvocationOutput = true;
    });
    builder.addCase(runFunction.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export const { setRequestBody, dismissInvocationOutput } = slice.actions;

export default slice.reducer;
