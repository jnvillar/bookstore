import axios from 'axios';

export const getBooks = () => {
  return axios.get(`/books`)
    .then(res => {
      return res.data
    })
}