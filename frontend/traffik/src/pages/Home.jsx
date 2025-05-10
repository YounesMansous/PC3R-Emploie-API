import React, { useEffect, useState } from "react";
import axios from "axios";
import { BASE_URL } from "../utils/const";
import { useNavigate } from "react-router";

function Home() {
  const [transportsMode, setTransportsMode] = useState([]);
  const navigate = useNavigate();

  const toTransportEventsHandler = (mode) => {
    console.log("hello");
    navigate(`/events/${mode}`);
  };

  useEffect(() => {
    const getTransportsMode = async () => {
      try {
        const response = await axios.get(`${BASE_URL}/lines/modes`);
        setTransportsMode(response.data.modes);
      } catch (error) {
        console.log("Error fetching transport modes:", error);
      }
    };

    getTransportsMode();
  }, []);

  return (
    <div className="container mt-5">
      <h1 className="text-center mb-3">Welcome to Traffik</h1>
      <p className="text-center">
        Your one-stop solution for traffic management.
      </p>
      <div className="row justify-content-center">
        {transportsMode.length === 0 ? (
          <p className="text-center">No transport modes available.</p>
        ) : (
          transportsMode.map((mode) => (
            <div
              onClick={() => {
                toTransportEventsHandler(mode.type);
              }}
              key={mode.type}
              className="col-md-3 m-2 p-3 border rounded text-center"
            >
              <h5>{mode.type}</h5>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default Home;
