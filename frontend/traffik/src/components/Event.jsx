import React from "react";

import "../styles/Events.css";
import { formToJSON } from "axios";

const Event = ({ id, title, message, date, ligne, onCommentClick }) => {
  return (
    <div className="event">
      <h5>Ligne {ligne.name}</h5>
      <hr />
      <h6 dangerouslySetInnerHTML={{ __html: title }}></h6>
      <div dangerouslySetInnerHTML={{ __html: message }}></div>
      <p className="date">{formToJSON(date)}</p>

      <button
        className="button-comments  align-self-center"
        data-bs-toggle="offcanvas"
        data-bs-target="#offcanvasRight"
        aria-controls="offcanvasRight"
        onClick={() => onCommentClick(id)}
      >
        Voir les commentaires
      </button>
    </div>
  );
};

export default Event;
