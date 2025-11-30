import React from 'react'
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCirclePlay } from '@fortawesome/free-solid-svg-icons'
import './Episode.css'

const Episode = ({ episode }) => {
  return (
    <div className="col-md-4 mb-4">
      <Link 
        to={`/stream/${episode.website_id}`} 
        style={{ textDecoration: "none", color: "inherit" }}
      >
      <div className="card h-100 shadow-sm show-card">
        <div style={{ position: "relative" }}>
          <img
            src={episode.episode_thumbnail} 
            alt={episode?.title ?? ""} 
            className="card-img-top"
            style={{
              objectFit: "contain",
              height: "250px",
              width: "100%",
            }}
          />
          <span className="play-icon-overlay">
                    <FontAwesomeIcon icon={faCirclePlay} />
                </span>
        </div>
        <div className="card-body d-flex flex-column">
          <h5 className="card-title">{episode?.title ?? ""} </h5>
          <p className="card-text mb-2">Episode #{episode.episode_number}</p>
          <p className="card-text mb-2">{episode.description}</p>
        </div>
      </div>
      </Link>
    </div>
  );
};

export default Episode;
