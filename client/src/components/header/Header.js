import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import React, { useState } from 'react';
import './style.css'
import { Contact } from "../contact/Contact";

export const Header = () => {
  const [showContact, setShow] = useState(false);

  const handleShowContact = (show) => {
    setShow(show);
  }

  return (
    <div>
      <Navbar bg="dark" expand="lg">
        <Navbar.Brand className={"brand"} href="">La librería</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav"/>
        <Navbar.Collapse className="justify-content-end">
          <Nav.Link className={"contact"} onClick={e => handleShowContact(true)}>Contacto</Nav.Link>
        </Navbar.Collapse>
      </Navbar>
      <Contact shouldShow={showContact} handleClose={e => handleShowContact(false)}/>
    </div>
  )
}