import './css/App.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route, useLocation } from "react-router-dom";
import { AnimatePresence, motion} from "framer-motion" 
import Home from "./pages/Home.jsx"
import About from "./pages/About.jsx"
import Movies from "./pages/Movies.jsx"
import Navbar from "./components/Navbar.jsx"

function AnimatedRoutes (){
  const location = useLocation();

  return (
    <AnimatePresence mode = "wait">
      <Routes location={location} key = {location.pathname}>
        <Route path="/" element= { <PageWrapper><Home /></PageWrapper>} />
        <Route path="/About" element= { <PageWrapper><About /></PageWrapper>} />
        <Route path="/Movies" element= { <PageWrapper><Movies /></PageWrapper>} />
      </Routes>
    </AnimatePresence>
  )
}


function PageWrapper({children}) {
  return (
    <motion.div
      initial={{ opacity: 0, y:20 }}
      animate={{ opacity: 1, y: 0}}
      exit={{ opacity: 0, y: 20 }}
      transition={{ duration: 0.4 }}
    > 
      {children}
    </motion.div>
  )
}

function App() {
  

  return (
    <>
    <BrowserRouter>
    <Navbar />
    <AnimatedRoutes />
    </BrowserRouter>
  
    </>
  )
}


export default App;
