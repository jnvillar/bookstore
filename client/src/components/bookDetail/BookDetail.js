import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import React, { useState } from 'react';
import './bookDetail.css'
import { formatPrice, getBookCode } from "../../lib/utils";
import { Contact } from "../contact/Contact";


export const BookDetail = ({book, onHide}) => {
  const [showContact, setShowContact] = useState(false)
  return (
    <div>
      <Modal show onHide={onHide} size="lg" aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header>
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
          <Button variant={"success"} onClick={() => setShowContact(true)}>Comprar</Button>
          <Button variant={"danger"} onClick={onHide}>Cerrar</Button>
        </Modal.Footer>
      </Modal>
      <Contact shouldShow={showContact} handleClose={() => setShowContact(false)}/>
    </div>
  );
}

export const BookInfo = ({title, value}) => {
  const show = value !== null
  return (
    show ? <div><b>{title}:</b> {value}</div> : null
  );
}