import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import createFunction from './createFunction';

interface FunctionFormState {
  id?: string;
  name: string;
  image: string;
  function?: Fn;
  skipLogging: boolean;
  errorMessage: string;
  showError?: boolean;
  loading?: boolean;
}

const initialState: FunctionFormState = {
  errorMessage: '',
  name: '',
  image: '',
  skipLogging: false,
};

const slice = createSlice({
  name: 'functionForm',
  initialState,
  reducers: {
    setName: (state, action: PayloadAction<string>) => {
      state.name = action.payload;
    },
    setImage: (state, action: PayloadAction<string>) => {
      state.image = action.payload;
    },
    setSkipLogging: (state, action: PayloadAction<boolean>) => {
      state.skipLogging = action.payload;
    },
  },
  extraReducers: builder => {
    builder.addCase(createFunction.pending, state => {
      state.loading = true;
    });
    builder.addCase(createFunction.fulfilled, (state, action) => {
      state.loading = false;
      state.function = action.payload as Fn;
    });
    builder.addCase(createFunction.rejected, (state, action) => {
      state.loading = false;
      state.showError = true;
      state.errorMessage = (action.payload as ErrorResponse).error;
    });
  },
});

export const { setName, setImage, setSkipLogging } = slice.actions;

export default slice.reducer;
