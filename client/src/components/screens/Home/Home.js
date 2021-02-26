import React, { useEffect, useState, useCallback } from 'react';
import Container from 'react-bootstrap/Container';
import { getBooks } from "../../../api/client";
import Form from 'react-bootstrap/Form';
import debounce from 'lodash.debounce';
import { Book } from "./Layout/Book";
import './style.css'

export const Home = () => {

  const [books, setBooks] = useState([])

  const debounceGetBooks = useCallback(
    debounce((search) => getBooks(search).then(r => {
      setBooks(r)
    }), 250), []
  );

  useEffect(() => {
    getBooks('').then(r => {
      setBooks(r)
    })
  }, []);

  const onSearchInput = (input) => {
    const search = input.target.value.toLowerCase()
    debounceGetBooks(search)
  }

  return (
    <div className={"page"}>
      <div className="header"/>
      <Container>
        <Form className="search">
          <Form.Group>
            <Form.Control type="text" placeholder="Buscar" onChange={onSearchInput}/>
          </Form.Group>
        </Form>
        <div className="books-container">
          {
            books.map(book => (
              <Book book={book}/>
            ))
          }
        </div>
      </Container>
    </div>
  )
}