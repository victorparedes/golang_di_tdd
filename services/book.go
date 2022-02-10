package services

import "errors"

/*
Esto funciona como constructor, crea el servicio, inyecta los libros productivos
y devuelve el objeto listo para usar.
*/
func NewService() Services {
	BookServices := BookServices{}
	BookServices.books = []Book{
		{
			Isbn:   "8435021521",
			Titulo: "Cita con rama",
		},
		{
			Isbn:   "0575077220",
			Titulo: "Rama II",
		},
		{
			Isbn:   "8498723779",
			Titulo: "El jardin de Rama",
		},
		{
			Isbn:   "8498729890",
			Titulo: "Rama relevada",
		},
	} // esta es la implementacion real de libros que inyectamos al servicio

	return BookServices
}

/*
Esta struct es la implementacion de la interface que esta descripta en el file services.go
*/
type BookServices struct {
	books []Book // Aqui estamos usando Inversion of Control ya que los libros los inyectamos por afuera
}

/*
Al agregar "(b BookServices)" delante del nombre de la funcion
estamos asignando esta funcion como un "metodo" de la sturcute BookServices.
Esto significa que tiene acceso a la propiedad "books" de la sturcute BookServices
y que ademas esta structure cumple con la interface "Services"
*/
func (b BookServices) SearchByIsbn(isbn string) (*Book, error) {
	for _, aBook := range b.books {
		if aBook.Isbn == isbn {
			return &aBook, nil
		}
	}

	return nil, errors.New("Book is not found")
}
