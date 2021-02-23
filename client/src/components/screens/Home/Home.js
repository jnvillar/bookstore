import React, { useEffect, useState } from 'react';
import Form from 'react-bootstrap/Form'
import Container from 'react-bootstrap/Container';
import './style.css'
import { getBooks } from "../../../api/client";
import { Book } from "./Layout/Book";

export const Home = () => {

  const [books, setBooks] = useState([])
  const [filteredBooks, setFilteredBooks] = useState([])

  useEffect(() => {
    getBooks().then(r => {
      setFilteredBooks(r)
      setBooks(r)
    })
  }, []);

  const onSearchInput = (input) => {
    const search = input.target.value.toLowerCase()

    if (search === '') {
      setFilteredBooks(books)
      return
    }

    setFilteredBooks(books.slice().filter(
      b => b.name.toLowerCase().includes(search)
    ))
  }

  return (
    <div>
      <div className="header"/>
      <Container>
        <Form className="search">
          <Form.Group controlId="formBasicEmail">
            <Form.Control type="text" placeholder="Buscar" onChange={onSearchInput}/>
          </Form.Group>
        </Form>
        <div className="books-container">
          {
            filteredBooks.map(book => (
              <Book book={book}/>
            ))
          }
        </div>
      </Container>
    </div>
  )
}