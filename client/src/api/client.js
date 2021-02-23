import axios from 'axios';

function getHost() {
  return 'http://localhost:8080'
}

export const getBooks = () => {
  return axios.get(`${getHost()}/books`)
    .then(res => {
      return res.data
    })
}