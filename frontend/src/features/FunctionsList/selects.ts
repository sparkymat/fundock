import { RootState } from '../../store';

export const selectFunctionsListLoading = (state: RootState): boolean =>
  state.functionsList.loading || false;
