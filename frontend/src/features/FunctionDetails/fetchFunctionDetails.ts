import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';
import Fn from '../../models/Fn';
import { ErrorResponse } from '../FunctionsList/fetchFunctions';

const fetchFunctionDetails = createAsyncThunk<Fn | ErrorResponse, string>(
  'features/fetchFunctions',
  async (functionName: string) => {
    try {
      const response = await axios.get(`/api/fn/${functionName}`);
      return response.data as Fn;
    } catch (error) {
      console.error(error);
      return { error: error };
    }
  },
);

export default fetchFunctionDetails;
