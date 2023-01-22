import Invocation from '../../models/Invocation';
import { RootState } from '../../store';

export const selectInvocationDetailsLoading = (state: RootState): boolean =>
  state.invocationDetails.loading || false;

export const selectInvocation = (state: RootState): Invocation | undefined =>
  state.invocationDetails.invocation;
