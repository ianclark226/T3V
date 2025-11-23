import Show from "../show/Show"

const Shows = ({shows, updateShowReview, message}) => {
  return (
    <div className="container mt-4">
        <div className="row">
            {shows && shows.length > 0 ? shows.map((show) => (
                <Show key={show._id} updateShowReview={updateShowReview} show={show}/>
            ))
            : <h2>{message}</h2>
        }
        </div>
    </div>
  )
}

export default Shows
