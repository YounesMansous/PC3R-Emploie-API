import React from "react";
import { NavLink, useNavigate } from "react-router";
import Button from "./Button";
import { useAuth } from "../hooks/AuthContext";
import "../styles/Button.css";

const Navbar = () => {
  const { isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();
    navigate("/login");
  };

  return (
    <nav className="navbar navbar-expand-lg bg-body-tertiary shadow-sm">
      <div className="container-fluid justify-content-between">
        <NavLink className="navbar-brand" to="/">
          <h5 className="text-info m-0">Traffik</h5>
        </NavLink>
        <div className="d-flex align-items-center">
          {isAuthenticated ? (
            <Button text="Logout" handler={handleLogout} color="button-lines" />
          ) : (
            <>
              <NavLink to="/register" className="button-lines mx-3 py-2 px-3">
                Sign In
              </NavLink>
              <NavLink to="/login" className="button-lines py-2 px-3">
                Login
              </NavLink>
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
