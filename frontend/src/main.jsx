import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import { Provider } from 'react-redux'
import './index.css'
import store from './store/store.js'
import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from 'react-router-dom'
import Leaderboard from './components/leaderboard/leaderboard.jsx'
import Game from './components/Game/Game.jsx'
import Layout from './Layout.jsx'

const route = createBrowserRouter(
  createRoutesFromElements(
    <Route path='/' element = {<Layout/>}>
    <Route path='/' element = {<App/>}/>
    <Route path='/game' element = {<Game/>}/>
    <Route path='/leaderboard' element = {<Leaderboard/>}/>
   
    </Route>
  )
)

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
  <Provider store={store}>
  <RouterProvider router={route}/>
  </Provider>
  </React.StrictMode>,
)
