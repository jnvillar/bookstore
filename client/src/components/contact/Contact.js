import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';


export const Contact = ({shouldShow, handleClose}) => {

  return (
    <Modal show={shouldShow} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>Contactate con nosotros</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <i className="fa fa-whatsapp whatsapp-icon"/>
        <b> Whastapp: </b>
        <a target={"_blank"} href={"https://wa.me/+5491149979027"}>
          1149979027
        </a>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="outline-dark" onClick={handleClose}>
          Cerrar
        </Button>
      </Modal.Footer>
    </Modal>
  );
}