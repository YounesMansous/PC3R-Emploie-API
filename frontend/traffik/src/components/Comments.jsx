import React, { useState, useEffect, useCallback } from "react";
import axios from "axios";
import "../styles/Events.css";
import "../styles/Comments.css";
import { BASE_URL, formatDate } from "../utils/const";
import { useAuth } from "../hooks/AuthContext";

const Form = ({ eventId, onCommentAdded }) => {
  const [text, setText] = useState("");
  const [message, setMessage] = useState("");

  const token = localStorage.getItem("authToken");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!text.trim()) return;

    try {
      await axios.post(
        `${BASE_URL}/comments/add?event_id=${eventId}`,
        {
          content: text,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setText("");
      onCommentAdded();
    } catch (error) {
      console.log(error.response.status);
      if (error.response.status === 401) {
        setMessage("Veuillez vous connecter");
      } else {
        setMessage("Impossible de se connecter au serveur.");
      }
    }
  };

  return (
    <form className="px-3" onSubmit={handleSubmit}>
      {message && <div className="alert alert-danger">{message}</div>}
      <label htmlFor="comment" className="form-label">
        Ajouter un commentaire
      </label>
      <textarea
        className="form-control"
        id="comment"
        rows="3"
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button className="button-comments mt-2 w-100" type="submit">
        Ajouter
      </button>
    </form>
  );
};

const Comment = ({ text, user, date }) => {
  return (
    <div className="comment">
      <p className="message">{text}</p>
      <div className="bottom">
        <p className="user">{user}</p>
        <p className="date">{formatDate(date)}</p>
      </div>
    </div>
  );
};

const Comments = ({ eventId }) => {
  const [comments, setComments] = useState([]);
  const { isAuthenticated } = useAuth();

  const fetchComments = useCallback(async () => {
    try {
      const response = await axios.get(
        `${BASE_URL}/comments?event_id=${eventId}`
      );
      setComments(response.data.comments);
    } catch (error) {
      console.error("Erreur chargement commentaires", error);
    }
  }, [eventId]);

  useEffect(() => {
    if (eventId) {
      fetchComments();
    }
  }, [eventId, fetchComments]);

  return (
    <div
      className="offcanvas offcanvas-end"
      tabIndex="-1"
      id="offcanvasRight"
      aria-labelledby="offcanvasRightLabel"
    >
      <div className="offcanvas-header">
        <h5 className="offcanvas-title" id="offcanvasRightLabel">
          Commentaires
        </h5>
        <button
          type="button"
          className="btn-close"
          data-bs-dismiss="offcanvas"
          aria-label="Close"
        ></button>
      </div>
      {isAuthenticated && (
        <Form eventId={eventId} onCommentAdded={fetchComments} />
      )}

      <div className="offcanvas-body bg-light opacity-3 mt-3 mx-3">
        {comments == null ? (
          <p>Aucun commentaire pour cet évènement.</p>
        ) : (
          comments.map((comment) => (
            <Comment
              key={comment.id}
              text={comment.content}
              user={comment.user}
              date={comment.created_at}
            />
          ))
        )}
      </div>
    </div>
  );
};

export default Comments;
