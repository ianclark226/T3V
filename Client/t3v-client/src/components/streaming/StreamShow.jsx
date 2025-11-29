import {useParams} from 'react-router-dom';
import ReactPlayer from 'react-player'
import './StreamShow.css'

const StreamShow = () => {

    let params = useParams();
    let key = params.website_id;

  console.log("website_id param:", params.website_id);

return (
    <div className="react-player-container">
      {(key!=null)?<ReactPlayer controls="true" playing={true} url ={`https://www.youtube.com/watch?v=${key}`} 
      width = '100%' height='100%' />:null}
    </div>
  )
}

export default StreamShow