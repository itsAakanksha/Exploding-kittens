import React from 'react'
import Navbar from './components/Common/Navbar'
import { Outlet } from 'react-router-dom'
function layout() {
  return (
    <>
    <Navbar/>
    <Outlet/>
    </>
  )
}

export default layout