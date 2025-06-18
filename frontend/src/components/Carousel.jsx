import {useState} from "react";  
import {Modal, Button} from "react-bootstrap"
import Slider from "react-slick";
import FilmCard from "./FilmCard";
import '../css/carousel.css'; // Import custom styles for the carousel

const MovieCarousel = ({ movies }) => {
  const sliderBehaviour = {
    dots: false,
    infinite: false,
    speed: 500,
    slidesToShow: 3,      // show 3 cards at once
    slidesToScroll: 1,    // scroll 1 card per arrow click
    arrows: true
    };

    const [selectedMovie, setSelectedMovie] = useState(null)

    const handleShowModal = (movie) => {
        setSelectedMovie(movie);
    };

    const handleCloseModal = () => {
        setSelectedMovie(null);
    };

  return (
    <div className="container">
      <Slider {...sliderBehaviour}>
        {movies.map((movie => (
          <div key={movie.FilmName}>
            <FilmCard onClick={() => handleShowModal(movie)}
              filmName={movie.FilmName}
              AverageRating={movie.AverageRating}
              totalVotes={movie.TotalVotes}
            />
          </div>
        )))}
      </Slider>
      {}
      {selectedMovie && (
        <Modal show={!!selectedMovie} onHide={handleCloseModal}>
        <Modal.Header closeButton>
            <Modal.Title>{selectedMovie.FilmName}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            <p>Rating: {selectedMovie.AverageRating}</p>
            <p>Votes: {selectedMovie.TotalVotes}</p>
        </Modal.Body>
        <Modal.Footer>
            <Button variant="secondary" onClick={handleCloseModal}>Close</Button>
        </Modal.Footer>
        </Modal>
      )}
    </div>
  );
};

export default MovieCarousel;
