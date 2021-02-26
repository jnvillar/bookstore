import React, { useEffect, useState, useCallback } from 'react';
import FormControl from 'react-bootstrap/FormControl';
import InputGroup from 'react-bootstrap/InputGroup';
import Container from 'react-bootstrap/Container';
import { getBooks } from "../../../api/client";
import ButtonToolbar from 'react-bootstrap/ButtonToolbar';
import Button from 'react-bootstrap/Button';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';
import Form from 'react-bootstrap/Form';
import debounce from 'lodash.debounce';
import { Book } from "../../book/Book";
import './home.css'

export const Home = () => {

  const priceTitle = {
    'asc': 'Mayor a menor',
    'desc': 'Menor a mayor',
    '': 'Precio'
  }

  const [books, setBooks] = useState([])
  const [advancedSearch, setAdvancedSearch] = useState(false)
  const [search, setSearch] = useState({search: '', price: ''})

  const debounceGetBooks = useCallback(
    debounce((search) => getBooks(search).then(r => {
      setBooks(r)
    }), 250), []
  );

  useEffect(() => {
    getBooks(search).then(r => {
      setBooks(r)
    })
  }, []);

  useEffect(() => {
    debounceGetBooks(search)
  }, [search]);


  const onAdvancedSearchClick = () => {
    setAdvancedSearch(!advancedSearch)
  }

  const onSetPriceOrder = (priceOrder) => {
    if (priceOrder === search['price']) {
      priceOrder = ''
    }
    setSearch({search: search['search'], price: priceOrder})
  }

  const onSearchInput = (input) => {
    const searchInput = input.target.value.toLowerCase()
    setSearch({search: searchInput, price: search['price']})
  }

  return (
    <div className={"page"}>
      <div className="header"/>
      <Container>
        <Form className="search">
          <Form.Group>
            <InputGroup className="mb-3">
              <FormControl onChange={onSearchInput} placeholder="Buscar"/>
              <InputGroup.Append>
                <Button onClick={onAdvancedSearchClick}
                        variant={advancedSearch ? "dark" : "outline-dark"}>Avanzado</Button>
              </InputGroup.Append>
            </InputGroup>
          </Form.Group>

          {advancedSearch
            ? <Form.Group>
              <ButtonToolbar aria-label="Toolbar with button groups">
                <ButtonGroup className="mr-2" aria-label="Second group">
                  <DropdownButton as={ButtonGroup} title={priceTitle[search['price']]} id="bg-vertical-dropdown-2"
                                  variant={"dark"}>
                    <Dropdown.Item active={search.price === 'desc'} onClick={e => onSetPriceOrder("desc")}>{priceTitle['desc']}</Dropdown.Item>
                    <Dropdown.Item active={search.price === 'asc' ? 'active' : ''} onClick={e => onSetPriceOrder("asc")}>{priceTitle['asc']}</Dropdown.Item>
                  </DropdownButton>
                </ButtonGroup>
              </ButtonToolbar>
            </Form.Group> : null
          }

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