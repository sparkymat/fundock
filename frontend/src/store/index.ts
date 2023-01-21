import { configureStore } from '@reduxjs/toolkit';
import functionsListReducer from '../features/FunctionsList/slice';
import functionDetailsReducer from '../features/FunctionDetails/slice';
import invocationsListReducer from '../features/InvocationsList/slice';
import invocationDetailsReducer from '../features/InvocationsDetails/slice';
import functionFormReducer from '../features/FunctionForm/slice';

export const store = configureStore({
  reducer: {
    functionsList: functionsListReducer,
    functionDetails: functionDetailsReducer,
    invocationsList: invocationsListReducer,
    invocationDetails: invocationDetailsReducer,
    functionForm: functionFormReducer,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
