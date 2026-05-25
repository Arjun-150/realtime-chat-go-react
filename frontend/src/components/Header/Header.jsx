import React, { useState } from "react";
import "./Header.css";

const Header = (props) => { // 🚀 Ensure props is passed here
  const [showModal, setShowModal] = useState(false);

  const handleClear = async () => {
    const host = window.location.hostname;
    try {
      const response = await fetch(`http://${host}:8080/clear`);
      if (response.ok) {
        setShowModal(false); // Close the popup
        props.onClear();     // 🚀 EXECUTE THE PROP: This clears the chatHistory in App.js
      }
    } catch (err) {
      console.error("Failed to clear chat:", err);
    }
  };
  return (
    <div className="header">
      {/* Invisible spacer to keep the title centered */}
      <div style={{ width: "80px" }}></div>

      <h2>REALTIME GROUPCHAT</h2>

      <button className="clear-btn" onClick={() => setShowModal(true)}>
        Clear
      </button>

      {showModal && (
        <div className="modal-overlay">
          <div className="modal">
            <h3>Wipe History?</h3>
            <p>This will delete everything from MongoDB permanently.</p>
            <div className="modal-actions">
              <button className="confirm-btn" onClick={handleClear}>Delete All</button>
              <button className="cancel-btn" onClick={() => setShowModal(false)}>Cancel</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default Header;