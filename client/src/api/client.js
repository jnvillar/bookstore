import axios from 'axios';

export const getBooks = (search) => {
  const url = `api/books?search=${search.search}&price=${search.price}`
  return axios.get(url)
    .then(res => {
      return res.data
    })
}