import axios from 'axios';

export const getBooks = (search) => {
  const url = `api/books?search=${search.search}&price=${search.price}&page=${search.page}&cat=${search.cat}`
  return axios.get(url)
    .then(res => {
      return res.data
    })
}

export const visitBook = (book) => {
  const url = `api/books/${book.id}/visit`
  return axios.get(url)
}

export const getCategories = () => {
  const url = `api/bookCategories`
  return axios.get(url)
    .then(res => {
      return res.data
    })
}