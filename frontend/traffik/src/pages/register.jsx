import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router";
import { useAuth } from "../hooks/AuthContext";
import { BASE_URL } from "../utils/const";
import "bootstrap/dist/css/bootstrap.min.css";

const Register = () => {
  const { login } = useAuth();
  const navigate = useNavigate();

  const [form, setForm] = useState({
    name: "",
    email: "",
    password: "",
  });

  const [isLoading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const validateForm = () => {
    if (!form.name.trim() || !form.email || !form.password) {
      setError("Tous les champs sont obligatoires.");
      return false;
    }
    if (!/\S+@\S+\.\S+/.test(form.email)) {
      setError("Email invalide.");
      return false;
    }
    if (form.password.length < 6) {
      setError("Le mot de passe doit contenir au moins 6 caractères.");
      return false;
    }
    return true;
  };

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    if (!validateForm()) return;

    try {
      setLoading(true);
      const res = await axios.post(`${BASE_URL}/register`, form);

      if (res.status === 200) {
        login();
        localStorage.setItem("authToken", res.data.token);
        navigate("/");
      }
    } catch (err) {
      console.error(err);
      if (err.response?.data?.error) {
        setError(err.response.data.error);
      } else {
        setError("Erreur lors de l'inscription.");
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="container d-flex justify-content-center align-items-center"
      style={{ minHeight: "100vh" }}
    >
      <div className="card p-4" style={{ maxWidth: "400px", width: "100%" }}>
        <h2 className="text-center mb-4">Inscription</h2>
        <form onSubmit={handleSubmit}>
          {error && <div className="alert alert-danger">{error}</div>}

          <div className="mb-3">
            <label htmlFor="name" className="form-label">
              Nom
            </label>
            <input
              type="text"
              id="name"
              name="name"
              value={form.name}
              onChange={handleChange}
              required
              className="form-control"
              disabled={isLoading}
            />
          </div>

          <div className="mb-3">
            <label htmlFor="email" className="form-label">
              Email
            </label>
            <input
              type="email"
              id="email"
              name="email"
              value={form.email}
              onChange={handleChange}
              required
              className="form-control"
              disabled={isLoading}
            />
          </div>

          <div className="mb-3">
            <label htmlFor="password" className="form-label">
              Mot de passe
            </label>
            <input
              type="password"
              id="password"
              name="password"
              value={form.password}
              onChange={handleChange}
              required
              className="form-control"
              disabled={isLoading}
            />
          </div>

          <button
            type="submit"
            className="btn btn-primary w-100"
            disabled={isLoading}
          >
            {isLoading ? "Création..." : "Créer un compte"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default Register;
