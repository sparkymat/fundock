import Fn from '../../models/Fn';
import Invocation from '../../models/Invocation';
import { RootState } from '../../store';

export const selectInvocationsListLoading = (state: RootState): boolean =>
  state.invocationsList.loading || false;

export const selectInvocations = (state: RootState): Invocation[] =>
  state.invocationsList.invocations;
