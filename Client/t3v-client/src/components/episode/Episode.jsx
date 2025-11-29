import React from 'react'

const Episode = ({ episode }) => {
  return (
    <div className="col-md-4 mb-4">
      <div className="card h-100 shadow-sm">
        <div style={{ position: "relative" }}>
          <img
            src={episode.episode_thumbnail} 
            alt={episode.title}
            className="card-img-top"
            style={{
              objectFit: "contain",
              height: "250px",
              width: "100%",
            }}
          />
        </div>
        <div className="card-body d-flex flex-column">
          <h5 className="card-title">{episode.title}</h5>
          <p className="card-text mb-2">Episode #{episode.episode_number}</p>
          <p className="card-text mb-2">{episode.description}</p>
        </div>
      </div>
    </div>
  );
};

export default Episode;
