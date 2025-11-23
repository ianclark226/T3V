import { useState } from 'react'
import './App.css'
import Home from './components/home/Home'
import Header from './components/header/Header'
import { Route, Routes, useNavigate } from 'react-router-dom'
import Register from './components/register/Register'
import Login from './components/login/Login'
import Layout from './components/Layout'
import RequiredAuth from './components/RequiredAuth'
import Recommended from './components/recommended/Recommended'
import Review from './components/review/Review'

function App() {

  const navigate = useNavigate()
  const updateShowReview = (show_id) => {
  navigate(`/review/${show_id}`)
}

  return (
    <>
    <Header />
      {/* <Header handleLogout = {handleLogout}/> */}
      <Routes path="/" element = {<Layout/>}>
        <Route path="/" element={<Home updateShowReview={updateShowReview}/>}></Route>
        <Route path="/register" element={<Register/>}></Route>
        <Route path="/login" element={<Login/>}></Route>
        <Route element = {<RequiredAuth/>}>
            <Route path="/recommended" element={<Recommended/>}></Route>
            {/* <Route path="/review/:imdb_id" element={<Review/>}></Route> */}
            {/* <Route path="/stream/:yt_id" element={<StreamMovie/>}></Route> */}
        </Route>
      </Routes>

    </>
  )
}

export default App
