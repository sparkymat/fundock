import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';
import createFunction from './createFunction';

interface KeyValuePair {
  key: string;
  value: string;
}

interface FunctionFormState {
  id?: string;
  name: string;
  image: string;
  environment: KeyValuePair[];
  secrets: KeyValuePair[];
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
  environment: [],
  secrets: [],
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
    addEnvironmentKey: state => {
      state.environment.push({ key: '', value: '' });
    },
    setEnvironmentValue: (
      state,
      action: PayloadAction<{ value: string; position: number }>,
    ) => {
      state.environment[action.payload.position].value = action.payload.value;
    },
    setEnvironmentKey: (
      state,
      action: PayloadAction<{ key: string; position: number }>,
    ) => {
      state.environment[action.payload.position].key = action.payload.key;
    },
    addSecretsKey: state => {
      state.secrets.push({ key: '', value: '' });
    },
    setSecretsValue: (
      state,
      action: PayloadAction<{ value: string; position: number }>,
    ) => {
      state.secrets[action.payload.position].value = action.payload.value;
    },
    setSecretsKey: (
      state,
      action: PayloadAction<{ key: string; position: number }>,
    ) => {
      state.secrets[action.payload.position].key = action.payload.key;
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

export const {
  setName,
  setImage,
  setSkipLogging,
  addEnvironmentKey,
  setEnvironmentValue,
  setEnvironmentKey,
  addSecretsKey,
  setSecretsKey,
  setSecretsValue,
} = slice.actions;

export default slice.reducer;
