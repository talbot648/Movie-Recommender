import InfoBox from '../components/InfoBox';
import Carousel from '../components/Carousel';
import '../css/Home.css'

 


const Home = () => {
      const movies = [
    { title: 'Inception', rating: 8.8, votes: 2000000 },
    { title: 'The Dark Knight', rating: 9.0, votes: 2500000 },
    { title: 'Interstellar', rating: 8.6, votes: 1800000 },
    { title: 'Tenet', rating: 7.4, votes: 400000 },
    { title: 'Dunkirk', rating: 7.9, votes: 600000 },
    { title: 'Memento', rating: 8.4, votes: 1200000 },
  ];

return (
<>
<div className="home-background">
<div className="home-container">
<h1>Find Movies You'll Love</h1>
<p>Discover Movies you love with the use of this Movie recommendation tool! </p>
</div>
</div>
<div className="grid-container">
 <InfoBox
    title = "About"
    description = "We recommend you movies based on your prerences and ratings."
    buttonText = "Learn More"
    buttonLink = "/About"
 />
  <InfoBox
    title = "Movies"
    description = "Find Your Next Favorite Movie:"
    buttonText = "Take Me There"
    buttonLink = "/Movies"
 />
</div>
<div className='container'>
    <h1>Top Rated Movies</h1>
    <Carousel movies={movies} />

</div>


</>
)}
export default Home;