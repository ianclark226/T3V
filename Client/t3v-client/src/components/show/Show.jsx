import Button from 'react-bootstrap/Button'

const Show = ({show, updateShowReview}) => {
  return (
    <div className="col-md-4 mb-4">
        <div className="card h-100 shadow-sm">
            <div style={{position:"relative"}}>
                <img src={show.poster_path} alt={show.title} 
                className="card-img-top"
                style={{
                    objectFit: "contain",
                    height: "250px",
                    width: "100%"
                }}
                />
            </div>
            <div className="card-body d-flex flex-column">
                <h5 className="card-title">{show.title}</h5>
                <p className="card-text mb-2">{show.show_id}</p>
            </div>
            {show.ranking?.ranking_name && (
                <span className="badge bg-dark m-3 p-2" style={{fontSize:"1rem"}}>
                    {show.ranking.ranking_name}
                </span>
            )}
            {updateShowReview && (
                <Button
                variant='outline-info'
                onClick={e => {
                    e.preventDefault();
                    updateShowReview(show.show_id)
                }}
                className='m-3'
                >
                    Review
                </Button>
            )}
        </div>
    </div>
  )
}

export default Show