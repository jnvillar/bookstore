import React from 'react';
import background from '../../../assets/books.jpg';
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import './style.css'

export const Home = () => {

  return (
    <div>
      <div className="header"/>
      <Container>
        <Form className="search">
          <Form.Group controlId="formBasicEmail">
            <Form.Control type="text" placeholder="Nombre del libro"/>
          </Form.Group>
        </Form>
        <div className="cardsContainer">
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
          <Card className="card">
            <Card.Img variant="top" src="https://i.pinimg.com/originals/7f/21/20/7f212010fb43e5d7cd839c5372e08433.jpg"/>
            <Card.Body>
              <Card.Title>Harry Potter</Card.Title>
              <Card.Text>
                Some quick example text to build on the Harry Potter and make up the bulk of
                the card's content.
              </Card.Text>
              <Button variant="primary">Go somewhere</Button>
            </Card.Body>
          </Card>
        </div>
      </Container>
    </div>
  )
}