import Button from 'react-bootstrap/Button'
import useAuth from '../../hook/useAuth'
import { Link, useNavigate } from 'react-router-dom'

const Show = ({show, updateShowReview}) => {

    const navigate = useNavigate()
    const { auth } = useAuth()

    const handleReviewClick = (e) => {
        if(!auth) {
            navigate('/login')
            return
        }

        updateShowReview(show.show_id)
    }

    const goToEpisodes = () => {
        navigate(`/shows/${show.show_id}/episodes`)
    }

  return (
    <div className="col-md-4 mb-4">
        <div className="card h-100 shadow-sm">
            <div style={{position:"relative"}}>
                <img src={show?.poster_path ?? ""} alt={show?.title ?? ""} 
                className="card-img-top"
                style={{
                    objectFit: "contain",
                    height: "250px",
                    width: "100%"
                }}
                />
            </div>
            <div className="card-body d-flex flex-column">
                <h5 className="card-title">{show?.title ?? ""} </h5>
            </div>
            {!!show?.ranking?.ranking_name && (
    <span className="badge bg-dark m-3 p-2" style={{fontSize:"1rem"}}>
        {show.ranking.ranking_name}
    </span>
)}
            {updateShowReview && (
                <Button
                variant='outline-info'
                onClick={handleReviewClick}
                className='m-3'
                >
                    Review
                </Button>
            )}
                    <h5 onClick={goToEpisodes} style={{ cursor: "pointer" }}>
                    </h5>
                

                <Button
                    className="btn btn-info m-3"
                    onClick={goToEpisodes}
                >
                    View Episodes
                </Button>
        </div>
    </div>
  )
}

export default Show