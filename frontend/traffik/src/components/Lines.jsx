import { useState } from "react";
import "../styles/Button.css";

const Line = ({ line, onClick }) => {
  return (
    <button className="animated-button w-100" onClick={() => onClick(line)}>
      {line.name}
    </button>
  );
};

const Lines = ({ lines, onSelect }) => {
  const [searchTerm, setSearchTerm] = useState("");

  const filteredLines = lines.filter((line) =>
    line.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div
      className="offcanvas offcanvas-start"
      data-bs-scroll="true"
      tabIndex="-1"
      id="offcanvasWithBothOptions"
      aria-labelledby="offcanvasWithBothOptionsLabel"
    >
      <div className="offcanvas-header">
        <h5 className="offcanvas-title" id="offcanvasWithBothOptionsLabel">
          List of lines
        </h5>
        <button
          type="button"
          className="btn-close"
          data-bs-dismiss="offcanvas"
          aria-label="Close"
        ></button>
      </div>

      <div className="px-3 pb-2">
        <input
          type="text"
          className="form-control"
          placeholder="Rechercher une ligne..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>

      <div className="offcanvas-body d-flex flex-column">
        {filteredLines.map((line) => (
          <Line key={line.id} line={line} onClick={onSelect} />
        ))}
        {filteredLines.length === 0 && <p>Aucune ligne trouvée</p>}
      </div>
    </div>
  );
};

export default Lines;
