import React from 'react';
import { Card } from 'react-bootstrap';

const FilmCard = ({ filmTitle, AverageRating, totalVotes }) => (
  <Card className="m-2 shadow-sm" style={{ minWidth: '250px' }}>
    <Card.Body className="text-center">
      <Card.Title>{filmTitle}</Card.Title>
      <Card.Text>⭐ {AverageRating} </Card.Text>
      <Card.Text>🗳 {totalVotes.toLocaleString()} votes</Card.Text>
    </Card.Body>
  </Card>
);

export default FilmCard;
