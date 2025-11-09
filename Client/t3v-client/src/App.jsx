import { useState } from 'react'
import './App.css'
import Home from './components/home/Home'
import Header from './components/header/Header'
import { Route, Routes, useNavigate } from 'react-router-dom'
import Register from './components/register/Register'
import Login from './components/login/Login'
import Layout from './components/Layout'
import RequiredAuth from './components/RequiredAuth'

function App() {

  return (
    <>
      <Header/>
      <Routes path="/" element={<Layout/>}>
        <Route path="/" element={<Home/>}/>
        <Route path="/register" element={<Register/>}/>
        <Route path="/login" element={<Login/>}/>
        <Route element={<RequiredAuth/>}>
        {/* <Route path="/recommended" element={<Recommended/>} /> */}
        </Route>
      </Routes>
    </>
  )
}

export default App
