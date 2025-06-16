import React from "react";
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

  return (
    <div className="container">
      <Slider {...sliderBehaviour}>
        {movies.map((movie, idx) => (
          <div key={idx}>
            <FilmCard
              filmTitle={movie.title}
              AverageRating={movie.rating}
              totalVotes={movie.votes}
            />
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default MovieCarousel;
