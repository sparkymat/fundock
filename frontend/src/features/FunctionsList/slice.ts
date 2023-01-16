import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface FunctionsListState {
  loading?: boolean;
}

const initialState: FunctionsListState = {};

const slice = createSlice({
  name: 'functionsList',
  initialState,
  reducers: {
    fetchFunctions: state => {
      state.loading = true;
    },
  },
});

export const { fetchFunctions } = slice.actions;

export default slice.reducer;
