import axios from 'axios';

export const getBooks = (search) => {
  return axios.get(`api/books?search=${search}`)
    .then(res => {
      return res.data
    })
}