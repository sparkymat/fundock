import Fn from '../../models/Fn';
import { RootState } from '../../store';

export const selectFunctionsListLoading = (state: RootState): boolean =>
  state.functionsList.loading || false;

export const selectFunctions = (state: RootState): Fn[] =>
  state.functionsList.functions;
