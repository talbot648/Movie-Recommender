import './css/App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home.jsx"
import About from "./pages/About.jsx"
import Navbar from "./components/Navbar.jsx"

function App() {
  

  return (
    <>
    <BrowserRouter>
    <Routes>
    <Route path="/" element={<Navbar />} >
    <Route index element={<Home />} />
    <Route path="about" element={<About />} />
    </Route>
    </Routes>
    </BrowserRouter>
  
    </>
  )
}

export default App;
