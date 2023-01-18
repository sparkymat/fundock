import Fn from '../../models/Fn';
import { RootState } from '../../store';

export const selectFunctionDetailsLoading = (state: RootState): boolean =>
  state.functionDetails.loading || false;

export const selectFunction = (state: RootState): Fn | undefined =>
  state.functionDetails.function;
