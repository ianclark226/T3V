import {Form, Button} from 'react-bootstrap';
import { useRef,useEffect,useState } from 'react';
import {useParams} from 'react-router-dom'
//import axiosPrivate from '../../api/axiosPrivateConfig';
import useAxiosPrivate from '../../hook/useAxiosPrivate';
import useAuth from '../../hook/useAuth';
import Show from '../show/Show';
import Spinner from '../spinner/Spinner';

const Review = () => {

    const [show, setShow] = useState({});
    const [loading, setLoading] = useState(false);
    const revText = useRef();
    const { show_id } = useParams();
    const {auth, setAuth} = useAuth();
    const axiosPrivate = useAxiosPrivate();

    useEffect(() => {
        const fetchShow = async () => {
            setLoading(true);
            try {
                const response = await axiosPrivate.get(`/show/${show_id}`);
                setShow(response.data);
                console.log(response.data);
            } catch (error) {
                console.error('Error fetching show:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchShow();

    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        setLoading(true);
        try {
            
            const response = await axiosPrivate.patch(`/update-review/${show_id}`, { admin_review: revText.current.value });
            console.log(response.data);           

            setShow(prev => ({
    ...prev,
    admin_review: response.data?.admin_review ?? prev.admin_review,
    ranking: response.data?.ranking_name
        ? {
            ranking_name: response.data.ranking_name,
            ranking_value: prev.ranking?.ranking_value ?? null
        }
        : prev.ranking ?? null
}));

        } catch (err) {
            console.error(err);
            if (err.response && err.response.status === 401) {
                 console.error('Unauthorized access - redirecting to login');
                 localStorage.removeItem('user');
                // setAuth(null);
            } else {
                console.error('Error updating review:', err);
            }

        } finally {
            setLoading(false);
        }
    }; 

    return (
      <>
        {loading ? (
            <Spinner />
        ) : (
            <div className="container py-5">
                <h2 className="text-center mb-4">Admin Review</h2>
                <div className="row justify-content-center">
                    <div className="col-12 col-md-6 d-flex align-items-center justify-content-center mb-4 mb-md-0">
                        <div className="w-100 shadow rounded p-3 bg-white d-flex justify-content-center align-items-center">
                            <Show show={show} />
                        </div>
                    </div>
                    <div className="col-12 col-md-6 d-flex align-items-stretch">
                        <div className="w-100 shadow rounded p-4 bg-light">
                            {auth && auth.role === "ADMIN" ? (
                                <Form onSubmit={handleSubmit}>
                                    <Form.Group className="mb-3" controlId="adminReviewTextarea">
                                        <Form.Label>Admin Review</Form.Label>
                                        <Form.Control
                                            ref={revText}
                                            required
                                            as="textarea"
                                            rows={8}
                                            defaultValue={show?.admin_review}
                                            placeholder="Write your review here..."
                                            style={{ resize: "vertical" }}
                                        />
                                    </Form.Group>
                                    <div className="d-flex justify-content-end">
                                        <Button variant="info" type="submit">
                                            Submit Review
                                        </Button>
                                    </div>
                                </Form> ):(
                                <div className="alert alert-info">{show.admin_review}</div>
                            )}                           
                        </div>
                    </div>
                </div>
            </div>
        )}
    </>      

    );
}

export default Review;
