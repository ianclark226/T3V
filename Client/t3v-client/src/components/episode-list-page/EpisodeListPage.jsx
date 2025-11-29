import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import axiosConfig from "../../api/axiosConfig";
import EpisodeList from "../episode-list/EpisodeList";

const EpisodeListPage = () => {
  const { show_id } = useParams();
  const [episodes, setEpisodes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [message, setMessage] = useState("");

  useEffect(() => {
    const fetchEpisodes = async () => {
      try {
        const response = await axiosConfig.get(`/shows/${show_id}/episodes`);
        setEpisodes(response.data);

        if (!response.data || response.data.length === 0) {
          setMessage("No episodes found.");
          }
      } catch (e) {
        console.error(e);
        setMessage("Error loading episodes");
      } finally {
        setLoading(false);
      }
    };

    fetchEpisodes();
  }, [show_id]);

  if (loading) return <h2>Loading episodes...</h2>;
  console.log("EPISODES RECEIVED:", episodes);

  return <EpisodeList episodes={episodes} message={message} />;
};

export default EpisodeListPage;