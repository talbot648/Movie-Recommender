import React from 'react';
import { Card } from 'react-bootstrap';
import '../css/FilmCard.css'

const FilmCard = ({ filmName, AverageRating, totalVotes, onClick }) => (
  <Card className="film-card m-2 shadow-sm" onClick={onClick}>
    <Card.Body className="text-center">
      <Card.Title>{filmName}</Card.Title>
      <Card.Text>⭐ {AverageRating} </Card.Text>
      <Card.Text>🗳 {totalVotes} votes</Card.Text>
    </Card.Body>
  </Card>
);

export default FilmCard;
