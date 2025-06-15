import { Button } from "bootstrap";
import { Link } from "react-router-dom";
import '../css/InfoBox.css';




const InfoBox = ({ title, description, buttonText, buttonLink}) => {
    return (
        <div className ="info-box">
            <h2>{title}</h2>
            <p>{description}</p>
            <Link to={buttonLink}>
            <button variant="primary">{buttonText}</button>
            </Link>
        </div>
    )
}
export default InfoBox;