import React, { createContext, useContext, useState } from "react";
import axios from "axios";
import { BASE_URL } from "../utils/const";

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const login = () => setIsAuthenticated(true);
  const logout = async () => {
    try {
      await axios.get(`${BASE_URL}/logout`, { withCredentials: true });
    } catch (e) {
      console.error("Erreur lors du logout :", e);
    }
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
