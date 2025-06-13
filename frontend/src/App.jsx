import { useState } from 'react'
import './css/App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home.jsx"
import About from "./pages/About.jsx"

function App() {
  

  return (
    <>
    <BrowserRouter>
    <Routes>
    <Route index element={<Home />} />
    <Route path="about" element={<About />} />
    </Routes>
    </BrowserRouter>
  
    </>
  )
}

export default App;
