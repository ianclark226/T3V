import useAxiosPrivate from "../../hook/useAxiosPrivate"
import { useEffect, useState } from "react"
import Shows from "../show/Show"

const Recommended = () => {

    const [shows, setShows] = useState([])
    const [loading, setLoading] = useState(false)
    const [message, setMessage] = useState()

    const axiosPrivate = useAxiosPrivate()

    useEffect(() => {
        const fetchRecommendedShows = async() => {
            setLoading(true)
            setMessage("")

            try {
                const response = await axiosPrivate.get('/recommended-shows')
                setShows(response.data)
                
            } catch (error) {
                console.log("Error fetching recommended shows:", error)
            } finally {
                setLoading(false)
            }
        }
        fetchRecommendedShows()
    }, [])
  return (
    <>
    {loading ? (
        <h2>Loading...</h2>
    ) : (
        <Shows shows={shows} message={message} />
    )}
    </>
  )
}

export default Recommended
