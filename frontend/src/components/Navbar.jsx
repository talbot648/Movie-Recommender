import { Outlet, NavLink } from "react-router-dom";
import { NavDropdown } from "react-bootstrap";
import '../css/navbar.css';
import { useState } from "react";

const Navbar = () => {
    const [isOpen, setIsOpen] = useState(false);

    
  return (
    <>
      <nav className="navbar navbar-expand-lg">
        <div className="container">
          <NavLink className="navbar-brand" to="/"></NavLink>

          <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" onClick={() => setIsOpen(!isOpen)} aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
    
          <div className={`collapse navbar-collapse ${isOpen ? 'show' : ''}`} id="navbarNav">

            <ul className="navbar-nav">
              <li className="nav-item">
                <NavLink className="nav-link" to="/">Home</NavLink>
              </li>
              <li className="nav-item">
                <NavLink className="nav-link" to="/About">About</NavLink>
              </li>
              <NavDropdown title="Movies">
               <li className="nav-item">
                <NavDropdown.Item><NavLink className="Nav-link" to="/Recommendations">Recommendations</NavLink></NavDropdown.Item>
                <NavDropdown.Item><NavLink className="Nav-link" to="/Browse">Browse All</NavLink></NavDropdown.Item>
                <NavDropdown.Item><NavLink className="Nav-link" to="/Favourites">Favourites</NavLink></NavDropdown.Item>
              </li>
              </NavDropdown>
            </ul>
            
          </div>
        </div>
      </nav>
      <div className="home-container">
        <Outlet />
      </div>
    </>
  )
};

export default Navbar;