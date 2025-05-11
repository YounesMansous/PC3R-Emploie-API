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
      const token = localStorage.getItem("authToken");

      await axios.post(
        `${BASE_URL}/logout`,
        {},
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
    } catch (e) {
      console.error("Erreur lors du logout :", e);
    } finally {
      localStorage.removeItem("authToken");
      setIsAuthenticated(false);
    }
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};
