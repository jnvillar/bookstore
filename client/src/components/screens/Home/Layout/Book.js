import React from 'react';
import Card from "react-bootstrap/Card";
import { formatPrice } from "../../../../lib/utils";

export const Book = ({book}) => (
  <Card className="card">
    <Card.Img variant="top" src={book.pictureUrl}/>
    <Card.Body>
      <Card.Text>{book.name}</Card.Text>
    </Card.Body>
    <Card.Footer>
      <div className={"price"}>
        <b>Precio:</b> {formatPrice(book.price)}
      </div>
    </Card.Footer>
  </Card>
);