import Fn from '../../models/Fn';
import Invocation from '../../models/Invocation';
import { RootState } from '../../store';

export const selectFunctionDetailsLoading = (state: RootState): boolean =>
  state.functionDetails.loading || false;

export const selectFunction = (state: RootState): Fn | undefined =>
  state.functionDetails.function;

export const selectRequestBody = (state: RootState): string =>
  state.functionDetails.requestBody;

export const selectInvocation = (state: RootState): Invocation | undefined =>
  state.functionDetails.invocation;

export const selectShowInvocationOutput = (state: RootState): boolean =>
  state.functionDetails.showInvocationOutput || false;
