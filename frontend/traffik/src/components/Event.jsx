import React from "react";

import "../styles/Events.css";

const formatDate = (isoDate) => {
  const dateObj = new Date(isoDate);
  return dateObj.toLocaleString("fr-FR", {
    day: "2-digit",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const Event = ({ id, title, message, date, ligne, onCommentClick }) => {
  return (
    <div className="event">
      <h5>Ligne {ligne.name}</h5>
      <hr />
      <h6 dangerouslySetInnerHTML={{ __html: title }}></h6>
      <div dangerouslySetInnerHTML={{ __html: message }}></div>
      <p className="date">{formatDate(date)}</p>

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
