import { useState, useEffect } from "react"
import axiosConfig from "../../api/axiosConfig"
import Shows from "../show/Show"

const Home = ({updateShowReview}) => {
    const [shows, setShows] = useState([])
    const [loading, setLoading] = useState(false)
    const [message, setMessage] = useState('')

    useEffect(() => {
        const fetchShows = async() => {
            setLoading(true)
            setMessage('')
            try {
                const response = await axiosConfig.get('/shows')
                setShows(response.data)
                if(response.data.length === 0) {
                    setMessage('There are currently no shows available')
                }
                
            } catch (error) {
                console.error('Error fetching shows:', error)
                setMessage('Error fetching shows')
            } finally {
                setLoading(false)
            }
        }
        fetchShows()
    }, [])
  return (
    <>
      {loading ? (
        <h2>Loading...</h2>
      ) : (
        <Shows shows={shows} message={message} updateShowReview={updateShowReview} />
      )}
    </>
  )
}

export default Home
