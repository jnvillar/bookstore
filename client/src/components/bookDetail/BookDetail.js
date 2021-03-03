import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import React from 'react';
import './bookDetail.css'
import { formatPrice, getBookCode } from "../../lib/utils";


export const BookDetail = ({book, onHide}) => {
  return (
    <Modal show onHide={onHide} size="lg" aria-labelledby="contained-modal-title-vcenter" centered>
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          {book.name}
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className={"book-container-modal"}>
          <img alt={book.name} src={book.pictureUrl} className={"book-image-detail"}/>
          <div className={"book-details-modal"}>
            <BookInfo title={"Autor"} value={book.author}/>
            <BookInfo title={"Precio"} value={formatPrice(book.price)}/>
            <BookInfo title={"Editorial"} value={book.publisher}/>
            <BookInfo title={"Código"} value={getBookCode(book)}/>
          </div>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant={"danger"} onClick={onHide}>Cerrar</Button>
      </Modal.Footer>
    </Modal>
  );
}

export const BookInfo = ({title, value}) => {
  const show = value !== null
  return (
    show ? <div><b>{title}:</b> {value}</div> : null
  );
}