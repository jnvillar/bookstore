import React, { useEffect, useState, useCallback } from 'react';
import FormControl from 'react-bootstrap/FormControl';
import InputGroup from 'react-bootstrap/InputGroup';
import { getBooks, getCategories, visitBook } from "../../../api/client";
import ButtonToolbar from 'react-bootstrap/ButtonToolbar';
import Button from 'react-bootstrap/Button';
import Alert from 'react-bootstrap/Alert';
import Spinner from 'react-bootstrap/Spinner';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';
import debounce from 'lodash.debounce';
import { Book } from "../../book/Book";
import './home.css'
import { BookDetail } from "../../bookDetail/BookDetail";
import { Page } from "../../pagination/Pagination";

export const Home = () => {

  const priceTitle = {
    'asc': 'Menor a mayor',
    'desc': 'Mayor a menor',
    '': 'Precio'
  }

  const [bookSelected, setBookSelected] = useState(null)
  const [categories, setCategories] = useState([])
  const [searching, setSearching] = useState(false)
  const [books, setBooks] = useState([])
  const [advancedSearch, setAdvancedSearch] = useState(false)
  const [search, setSearch] = useState({search: '', price: '', page: 0, cat: ''})

  const onSelectBook = (book) => {
    if (book !== null) {
      visitBook(book)
    }
    setBookSelected(book)
  }

  const onPageChange = (page) => {
    const newPage = Math.max(0, search.page + page)
    setSearch({search: search['search'], price: search.price, page: newPage, cat: search.cat})
  }

  const debounceGetBooks = useCallback(
    debounce((search) => getBooks(search).then(r => {
      setBooks(r)
      setSearching(false)
    }), 200), []
  );

  useEffect(() => {
    getBooks(search).then(r => {
      setBooks(r)
    })
    getCategories().then(r => {
      setCategories(r)
    })
  }, []);

  useEffect(() => {
    setSearching(true)
    debounceGetBooks(search)
  }, [search]);

  const onAdvancedSearchClick = () => {
    setAdvancedSearch(!advancedSearch)
  }

  const onSetPriceOrder = (priceOrder) => {
    if (priceOrder === search['price']) {
      priceOrder = ''
    }
    setSearch({search: search.search, price: priceOrder, page: search.page, cat: search.cat})
  }

  const onSetCategory = (category) => {
    if (category === search.cat) {
      category = ''
    }
    setSearch({search: search.search, price: search.price, page: 0, cat: category})
  }

  const onSearchInput = (input) => {
    const searchInput = input.target.value.toLowerCase()
    setSearch({search: searchInput, price: search.price, page: 0, cat: search.cat})
  }

  return (
    <div className={"page"}>
      <div className="header"/>
      <div className={"page-container"}>

        <div className={"search"}>
          <InputGroup className="mb-3">
            <FormControl onChange={onSearchInput} placeholder="Buscar"/>
            <InputGroup.Append>
              <Button onClick={onAdvancedSearchClick}
                      variant={advancedSearch ? "dark" : "outline-dark"}>Avanzado</Button>
            </InputGroup.Append>
          </InputGroup>

          {advancedSearch
            ?
            <ButtonToolbar aria-label="Toolbar with button groups">
              <ButtonGroup className="mr-2" aria-label="Second group">
                <DropdownButton as={ButtonGroup} title={priceTitle[search.price]} id="bg-vertical-dropdown-2"
                                variant={"dark"}>
                  <Dropdown.Item active={search.price === 'desc'}
                                 onClick={() => onSetPriceOrder("desc")}>{priceTitle['desc']}</Dropdown.Item>
                  <Dropdown.Item active={search.price === 'asc' ? 'active' : ''}
                                 onClick={() => onSetPriceOrder("asc")}>{priceTitle['asc']}</Dropdown.Item>
                </DropdownButton>
              </ButtonGroup>

              <ButtonGroup className="mr-2" aria-label="Second group">
                <DropdownButton as={ButtonGroup}
                                title={search.cat === '' ? "Categoría" : search.cat.toLowerCase().replace(/\b(\w)/g, s => s.toUpperCase())}
                                id="bg-vertical-dropdown-2"
                                variant={"dark"}>
                  {categories.map(cat =>
                    <Dropdown.Item active={search.cat === cat}
                                   onClick={() => onSetCategory(cat)}>{cat.toLowerCase().replace(/\b(\w)/g, s => s.toUpperCase())}</Dropdown.Item>
                  )}
                </DropdownButton>
              </ButtonGroup>
            </ButtonToolbar> : null
          }
        </div>

        {bookSelected
          ? <BookDetail book={bookSelected} onHide={() => onSelectBook(null)}/> : null
        }

        {searching
          ? <Spinner animation="border" role="status" className={"loading"}>
            <span className="sr-only">Loading...</span>
          </Spinner> : null
        }

        {!searching && books.length > 0
          ?
          <div className={"results"}>
            <Page currentPage={search.page} onPageChange={onPageChange}/>
            <div className="books-container">
              {
                books.map(book => (
                  <Book book={book} selectBook={() => onSelectBook(book)}/>
                ))
              }</div>
            <Page currentPage={search.page} onPageChange={onPageChange}/>
          </div> : null
        }

        {!searching && books.length == 0
          ? <Alert variant={"warning"}>
            No hay resultados
          </Alert> : null
        }
      </div>
    </div>
  )
}