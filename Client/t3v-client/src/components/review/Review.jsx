import { Form, Button } from 'react-bootstrap';
import { useRef, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import useAxiosPrivate from '../../hook/useAxiosPrivate';
import useAuth from '../../hook/useAuth';
import Show from '../show/Show';
import Spinner from '../spinner/Spinner';

const Review = () => {
    const [show, setShow] = useState(null);     // <-- null is safer than {}
    const [loading, setLoading] = useState(false);
    const revText = useRef();
    const { show_id } = useParams();
    const { auth } = useAuth();
    const axiosPrivate = useAxiosPrivate();

    useEffect(() => {
        const fetchShow = async () => {
            setLoading(true);
            try {
                const response = await axiosPrivate.get(`/show/${show_id}`);
                setShow(response.data);
                console.log("Show loaded:", response.data);
            } catch (error) {
                console.error('Error fetching show:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchShow();
    }, [show_id]);   // <-- important fix

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        try {
            const response = await axiosPrivate.patch(
                `/update-review/${show_id}`,
                { admin_review: revText.current.value }
            );

            console.log("Review updated:", response.data);

            // Safe ranking update
            setShow(prev => ({
                ...prev,
                admin_review: response.data?.admin_review ?? prev?.admin_review,
                ranking: response.data?.ranking_name
                    ? {
                        ranking_name: response.data.ranking_name,
                        ranking_value: prev?.ranking?.ranking_value ?? null
                    }
                    : prev?.ranking ?? null
            }));

        } catch (err) {
            console.error("Review update error:", err);
        } finally {
            setLoading(false);
        }
    };

    // 🔥 SAFETY CHECKS — prevents "undefined show_id" errors
    if (loading) return <Spinner />;
    if (!show) return <h3 className="text-center mt-5">Loading show...</h3>;
    if (!show?.show_id) return <h3 className="text-center mt-5">Invalid show</h3>;

    return (
        <div className="container py-5">
            <h2 className="text-center mb-4">Admin Review</h2>

            <div className="row justify-content-center">
                
                {/* Left Column — Show Preview */}
                <div className="col-12 col-md-6 mb-4">
                    <div className="shadow rounded p-3 bg-white">
                        <Show show={show} />
                    </div>
                </div>

                {/* Right Column — Review Form */}
                <div className="col-12 col-md-6">
                    <div className="shadow rounded p-4 bg-light">

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
                            </Form>

                        ) : (
                            <div className="alert alert-info">
                                {show?.admin_review ?? "No admin review available."}
                            </div>
                        )}

                    </div>
                </div>
            </div>
        </div>
    );
};

export default Review;