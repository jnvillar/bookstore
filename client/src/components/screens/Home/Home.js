import React, { useEffect, useState, useCallback } from 'react';
import Container from 'react-bootstrap/Container';
import { getBooks } from "../../../api/client";
import Form from 'react-bootstrap/Form';
import debounce from 'lodash.debounce';
import { Book } from "./Layout/Book";
import axios from 'axios';
import './style.css'

export const Home = () => {

  const [books, setBooks] = useState([])

  const debounceGetBooks = useCallback(
    debounce((search) => getBooks(search).then(r => {
      setBooks(r)
    }), 300), []
  );

  useEffect(() => {
    getBooks('').then(r => {
      setBooks(r)
      let config = {headers: {
        'Access-Control-Allow-Origin': '*'  //the token is a variable which holds the token
      }}
      axios.get(r[0]['pictureUrl'], config).then(r => console.log(r.data))
    })
  }, []);

  const onSearchInput = (input) => {
    const search = input.target.value.toLowerCase()
    debounceGetBooks(search)
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
            books.map(book => (
              <Book book={book}/>
            ))
          }
        </div>
      </Container>
    </div>
  )
}