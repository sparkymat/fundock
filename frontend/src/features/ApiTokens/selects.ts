import ApiToken from '../../models/ApiToken';
import Fn from '../../models/Fn';
import { RootState } from '../../store';

export const selectApiTokensLoading = (state: RootState): boolean =>
  state.apiTokens.loading || false;

export const selectApiTokens = (state: RootState): ApiToken[] =>
  state.apiTokens.apiTokens;

export const selectNewClientName = (state: RootState): string =>
  state.apiTokens.newClientName;
