import { Outlet, NavLink } from "react-router-dom";
import '../css/navbar.css';

const Navbar = () => {
  return (
    <>
      <nav className="navbar navbar-expand-lg">
        <div className="container">
          <NavLink className="navbar-brand" to="/"></NavLink>
          <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item">
                <NavLink className="nav-link" to="/">Home</NavLink>
              </li>
              <li className="nav-item">
                <NavLink className="nav-link" to="/About">About</NavLink>
              </li>
               <li className="nav-item">
                <NavLink className="nav-link" to="/Movies">Movies</NavLink>
              </li>
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