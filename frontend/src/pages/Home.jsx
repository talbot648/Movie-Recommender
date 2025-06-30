import InfoBox from '../components/InfoBox';
import Carousel from '../components/Carousel';
import Footer from '../components/Footer';
import '../css/Home.css'
import '../css/footer.css'
import { useEffect, useState } from 'react';

 


const Home = () => {
    const [Movies, setMovies] = useState([]);


    useEffect(() => {
    const fetchTopMovies = async () => {
        try {
            
            const response = await fetch("api/topMovies");
        
        if (!response.ok){
            throw new Error('Failed to fetch data')
        }

        const data = await response.json();
        setMovies(data)
    }   catch (error) {
        console.error('Error fetching the top movies:', error)
    }
    };

    fetchTopMovies();
    },[]);

console.log("Movies state in Home:", Movies);

return (
<>
<div className="home-background">
<div className="home-container">
<h1>Find Movies You'll Love</h1>
<p>Discover movies you love with the use of this recommendation tool! </p>
</div>
</div>
<div className="grid-container box">
 <InfoBox
    title = "About"
    description = "We recommend you movies based on your prerences and ratings."
    buttonText = "Learn More"
    buttonLink = "/About"
 />
  <InfoBox
    title = "Movies"
    description = "Find your next favorite movie:"
    buttonText = "Take Me There"
    buttonLink = "/Movies"
 />
</div>
<div className='container'>
    <h1>Top Rated Movies</h1>
    <Carousel movies={Movies} />

</div>

<Footer></Footer>



</>
)}
export default Home;