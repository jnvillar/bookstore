import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import React from 'react';
import './header.css'


export const Header = ({showContact}) => {
  return (
    <div>
      <Navbar bg="dark" expand="lg" sticky="top">
        <Navbar.Brand className={"brand"} href="">La librería</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav"/>
        <Navbar.Collapse bg="light" className="justify-content-end">
          <Nav.Link className={"contact"} onClick={() => showContact(true)}>Contacto</Nav.Link>
        </Navbar.Collapse>
      </Navbar>
    </div>
  )
}