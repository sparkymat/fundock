import Fn from '../../models/Fn';
import { RootState } from '../../store';

export const selectName = (state: RootState): string => state.functionForm.name;

export const selectImage = (state: RootState): string =>
  state.functionForm.image;

export const selectSkipLogging = (state: RootState): boolean =>
  state.functionForm.skipLogging;

export const selectFunction = (state: RootState): Fn | undefined =>
  state.functionForm.function;

export const selectFormProcessing = (state: RootState): boolean =>
  state.functionForm.loading || false;
