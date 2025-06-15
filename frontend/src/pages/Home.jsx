import InfoBox from '../components/InfoBox';
import '../css/Home.css'
 


const Home = () => {

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
    description = "Find Your Next Favorite Movie"
    buttonText = "Take Me There"
    buttonLink = "/Movies"
 />
</div>


</>
)}
export default Home;