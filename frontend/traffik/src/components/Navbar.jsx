import React from "react";
import { NavLink, useNavigate } from "react-router";
import Button from "./Button";
import { useAuth } from "../hooks/AuthContext";

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
        <div className="d-flex">
          {isAuthenticated ? (
            <Button
              text="Logout"
              handler={handleLogout}
              color="btn btn-light"
            />
          ) : (
            <>
              <Button
                text="Sign In"
                handler={() => console.log("sign in")}
                color="btn btn-light"
              />
              <NavLink to="/login" className="btn btn-light">
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
