import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { BASE_URL } from "../utils/const";
import axios from "axios";
import Lines from "../components/Lines";
import "../styles/Button.css";
import Event from "../components/Event";
import Comments from "../components/Comments";

const Events = () => {
  const params = useParams();

  const [lines, setLines] = useState([]);
  const [line, setLine] = useState({});
  const [events, setEvents] = useState([]);
  const [selectedEventId, setSelectedEventId] = useState(null);

  const [currentPage, setCurrentPage] = useState(1);
  const eventsPerPage = 5;

  const indexOfLastEvent = currentPage * eventsPerPage;
  const indexOfFirstEvent = indexOfLastEvent - eventsPerPage;
  const currentEvents = events.slice(indexOfFirstEvent, indexOfLastEvent);

  const totalPages = Math.ceil(events.length / eventsPerPage);

  useEffect(() => {
    const getTransportsModeLines = async () => {
      try {
        const response = await axios.get(`${BASE_URL}/lines/modes/id`, {
          params: { mode: params.mode },
        });
        const allLines = response.data.lines;
        setLines(allLines);
        if (allLines != null) {
          if (allLines.length > 0) {
            setLine(allLines[0]);
          }
        }
      } catch (error) {
        console.log("Error fetching transport modes:", error);
      }
    };

    getTransportsModeLines();
  }, [params.mode]);

  useEffect(() => {
    const getEvents = async () => {
      if (!line) return;
      try {
        const response = await axios.get(`${BASE_URL}/events/line`, {
          params: { id_line: line.id },
        });

        if (response.data.events != null) {
          setEvents(response.data.events);
        } else {
          setEvents([]);
        }
      } catch (error) {
        console.log("Error fetching events:", error);
      }
    };

    getEvents();
  }, [line]);

  useEffect(() => {
    setCurrentPage(1);
  }, [events]);

  const handleLineClick = (line) => {
    setLine(line);
  };

  return (
    <div>
      <div className="d-flex p-3 w-100">
        <button
          class="button-lines"
          type="button"
          data-bs-toggle="offcanvas"
          data-bs-target="#offcanvasWithBothOptions"
          aria-controls="offcanvasWithBothOptions"
        >
          {params.mode} lines
        </button>
      </div>

      <div className="container-fluid row mt-5 ">
        <Lines lines={lines} onSelect={handleLineClick} />
        <div className="mt-4 row justify-content-center align-items-center ">
          <h4 className="text-center">Évènements pour la ligne {line.name}</h4>
          {events.length === 0 ? (
            <p className="text-center mt-4">Aucun événement</p>
          ) : (
            <div className="col-sm-12 col-lg-5">
              {currentEvents.map((event) => (
                <Event
                  key={event.id}
                  id={event.id}
                  title={event.titre}
                  message={event.message}
                  date={event.created_at}
                  ligne={line}
                  onCommentClick={setSelectedEventId}
                />
              ))}
            </div>
          )}
          {events.length > eventsPerPage && (
            <div className="d-flex justify-content-center mt-3">
              <nav>
                <ul className="pagination">
                  <li
                    className={`page-item ${
                      currentPage === 1 ? "disabled" : ""
                    }`}
                  >
                    <button
                      className="page-link"
                      onClick={() => setCurrentPage(currentPage - 1)}
                    >
                      Précédent
                    </button>
                  </li>

                  {Array.from({ length: totalPages }, (_, i) => (
                    <li
                      key={i}
                      className={`page-item ${
                        currentPage === i + 1 ? "active" : ""
                      }`}
                    >
                      <button
                        className="page-link"
                        onClick={() => setCurrentPage(i + 1)}
                      >
                        {i + 1}
                      </button>
                    </li>
                  ))}

                  <li
                    className={`page-item ${
                      currentPage === totalPages ? "disabled" : ""
                    }`}
                  >
                    <button
                      className="page-link"
                      onClick={() => setCurrentPage(currentPage + 1)}
                    >
                      Suivant
                    </button>
                  </li>
                </ul>
              </nav>
            </div>
          )}
        </div>
        <Comments eventId={selectedEventId} />
      </div>
    </div>
  );
};

export default Events;
