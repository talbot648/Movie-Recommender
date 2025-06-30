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
    slidesToShow: 3,
    slidesToScroll: 1,
    arrows: true,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
          slidesToScroll: 1,
          infinite: true,
          dots: true,
        },
      },
      {
        breakpoint: 768,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
        },
      },
    ],
  };

    const [selectedMovieDetails, setSelectedMovieDetails] = useState(null);

    const handleShowModal = async (movie) => {
        try {
            const response = await fetch (`http://localhost:8080/api/movieDetails/${movie.Filmid}`);

            if (!response.ok){
                throw new Error('Failed to fetch data') 
            }
            
            const data = await response.json();
            setSelectedMovieDetails(data);
            }  catch (error) {
            console.error('Error fetching the movies details:', error)
            setSelectedMovieDetails(null);
        }
    };
        

    const handleCloseModal = () => {
        setSelectedMovieDetails(null);
    };

  return (
    <div className="container">
      <Slider {...sliderBehaviour}>
        {movies.map((movie => (
          <div key={movie.FilmId}>
            <FilmCard onClick={() => handleShowModal(movie)}
              filmName={movie.FilmName}
              AverageRating={movie.AverageRating}
              totalVotes={movie.TotalVotes}
            />
          </div>
        )))}
      </Slider>
      
      {selectedMovieDetails && (
        <Modal show={!!selectedMovieDetails} onHide={handleCloseModal}>
        <Modal.Header closeButton>
        <Modal.Title>{selectedMovieDetails.FilmName}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            <p>Overview: {selectedMovieDetails.Overview}</p>
            <p>IsAdult: {selectedMovieDetails.Adult.toString()}</p>
            <p>Rating: {selectedMovieDetails.AverageRating}</p>
            <p>Votes: {selectedMovieDetails.TotalVotes}</p>
            <p>Language: {selectedMovieDetails.Language}</p>
            <p>Released: {new Date(selectedMovieDetails.ReleaseDate).toLocaleDateString()}</p>
            
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
