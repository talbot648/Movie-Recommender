import React from 'react';
import { Card } from 'react-bootstrap';

const FilmCard = ({ filmName, AverageRating, totalVotes }) => (
  <Card className="m-2 shadow-sm" style={{ minWidth: '250px', fontfamily: 'Poppins, sans-serif' }}>
    <Card.Body className="text-center">
      <Card.Title>{filmName}</Card.Title>
      <Card.Text>⭐ {AverageRating} </Card.Text>
      <Card.Text>🗳 {totalVotes} votes</Card.Text>
    </Card.Body>
  </Card>
);

export default FilmCard;
