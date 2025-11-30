import Button from 'react-bootstrap/Button'
import useAuth from '../../hook/useAuth'
import { Link, useNavigate } from 'react-router-dom'

const Show = ({ show = {}, updateShowReview }) => {
    const navigate = useNavigate();
    const { auth } = useAuth();

    const safeId = show?.show_id;      // ← new
    const safeTitle = show?.title ?? "";
    const safePoster = show?.poster_path ?? "";

    const handleReviewClick = () => {
        if (!auth) {
            navigate('/login');
            return;
        }
        if (!safeId) return; // ← important guard
        updateShowReview(safeId);
    };

    const goToEpisodes = () => {
        if (!safeId) return; // ← important guard
        navigate(`/shows/${safeId}/episodes`);
    };

    return (
        <div className="col-md-4 mb-4">
            <div className="card h-100 shadow-sm">

                <div style={{ position:"relative" }}>
                    <img
                        src={safePoster}
                        alt={safeTitle}
                        className="card-img-top"
                        style={{
                            objectFit: "contain",
                            height: "250px",
                            width: "100%"
                        }}
                    />
                </div>

                <div className="card-body d-flex flex-column">
                    <h5 className="card-title">{safeTitle}</h5>
                </div>

                {!!show?.ranking?.ranking_name && (
                    <span className="badge bg-dark m-3 p-2" style={{ fontSize:"1rem" }}>
                        {show.ranking.ranking_name}
                    </span>
                )}

                {updateShowReview && (
                    <Button
                        variant="outline-info"
                        onClick={handleReviewClick}
                        className="m-3"
                    >
                        Review
                    </Button>
                )}

                <Button
                    className="btn btn-info m-3"
                    onClick={goToEpisodes}
                    disabled={!safeId}   // ← disable until loaded
                >
                    View Episodes
                </Button>

            </div>
        </div>
    );
};

export default Show