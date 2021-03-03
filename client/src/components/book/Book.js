import React from 'react';
import Card from "react-bootstrap/Card";
import { formatPrice } from "../../lib/utils";
import './book.css'

export const Book = ({book, selectBook}) => {
  return (
    <Card className={'book'} onClick={() => selectBook(book)}>
      <Card.Img className={'book-image'} variant="top" src={book.pictureUrl}/>
      <div className={'book-body'}>
        <Card.Body>
          <Card.Text>{book.name}</Card.Text>
        </Card.Body>
        <Card.Footer className={'book-footer'}>
          <div className={'book-price'}>
            <b>Precio:</b> {formatPrice(book.price)}
          </div>
        </Card.Footer>
      </div>
    </Card>
  );
}