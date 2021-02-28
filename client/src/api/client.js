import axios from 'axios';

export const getBooks = (search) => {
  const url = `api/books?search=${search.search}&price=${search.price}&page=${search.page}&cat=${search.cat}`
  console.log(url)
  return axios.get(url)
    .then(res => {
      return res.data
    })
}

export const getCategories = () => {
  const url = `api/books/categories`
  return axios.get(url)
    .then(res => {
      return res.data
    })
}