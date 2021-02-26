import React from 'react';
import Card from "react-bootstrap/Card";
import { formatPrice } from "../../lib/utils";
import { isMobile } from 'react-device-detect';
import './book.css'

export const Book = ({book}) => {
  return (
    <Card className={'book'}>
      <Card.Img className={isMobile ? 'book-image-mobile' : 'book-image'} variant="top" src={book.pictureUrl}/>
      <Card.Body className={isMobile ? 'book-body-mobile' : 'book-body'}>
        <Card.Text>{book.name}</Card.Text>
      </Card.Body>
      <Card.Footer>
        <div className={isMobile ? 'book-price-mobile' : 'book-price'}>
          <b>Precio:</b> {formatPrice(book.price)}
        </div>
      </Card.Footer>
    </Card>
  );
}