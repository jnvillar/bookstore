import Pagination from "react-bootstrap/Pagination";
import React from "react";
import './pagination.css'

export const Page = ({currentPage, onPageChange}) => {
  return (<Pagination className={"pagination"}>
    <Pagination.Prev onClick={() => onPageChange(-1)}/>
    <Pagination.Item> Página: {currentPage + 1}</Pagination.Item>
    <Pagination.Next onClick={() => onPageChange(1)}/>
  </Pagination>)
}
