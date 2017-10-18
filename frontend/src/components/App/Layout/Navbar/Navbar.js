import React from 'react';
import { NavLink } from 'react-router-dom';

const Navbar = () => (
  <nav className="navbar navbar-expand-md navbar-dark bg-dark">
    <NavLink exact to='/' className="navbar-brand">Charts App</NavLink>
    <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#meanAuthNavbar" aria-controls="meanAuthNavbar"
      aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
    </button>
    <div className="collapse navbar-collapse" id="meanAuthNavbar">
      <ul className="navbar-nav mr-auto">
        <li className="nav-item">
          <NavLink exact to='/' className="nav-link" activeClassName="active">Home</NavLink>
        </li>
      </ul>
    </div>
  </nav>
);

export default Navbar;