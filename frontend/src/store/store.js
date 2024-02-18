// src/app/store.js
import { configureStore } from '@reduxjs/toolkit';
import gameSlice from '../features/gameSlice';
// import userReducer from '../features/user/userSlice';

const store = configureStore({
  reducer: {
    game: gameSlice,
    // user: userReducer,
  },
});

export default store;
