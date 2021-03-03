export const formatPrice = (price) => {
  return `$${price / 100}`
}

export const getBookCode = (book) => {
  return `${(book.pictureUrl.includes('mlstatic') ? "m" : "b") + book.id.substring(0, 3)}`
}