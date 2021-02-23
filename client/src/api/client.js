import axios from 'axios';

export const getBooks = () => {
  return axios.get(`api/books`)
    .then(res => {
      return res.data
    })
}