import React from 'react'
import Episode from '../episode/Episode'

const EpisodeList = ({ episodes, message }) => {
  return (
    <div className="container mt-4">
        <div className="row">
            {episodes && episodes.length > 0 ? (
          episodes.map((episode) => (
            <Episode key={episode._id || episode.episode_id} episode={episode} />
            ))
        ) : (
          <h2>{message}</h2>
        )}
        </div>
    </div>
  )
}

export default EpisodeList
