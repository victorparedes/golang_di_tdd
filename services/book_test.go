package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Todos los test van a tener la misma estructura
   Arrange-Preparacion : creamos los datos falsos y el objeto a probar
   Act-Ejecucion       : Realizamos la llamada a la funcion que queremos testear.
   Assert-Verificacion : Comprobamos que lo que trajo es lo que esperabamos
*/

func TestBookServices(t *testing.T) {
	t.Run("When exist isbn must return a book", func(t *testing.T) {
		// Arrange
		target := BookServices{}
		target.books = []Book{
			{
				Isbn:   "un isbn",
				Titulo: "Un libro mockeado",
			},
		} // Esta es la implementacion "fake" de los libros y que me sirve solo para este test

		// Act
		result, err := target.SearchByIsbn("un isbn")

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "Un libro mockeado", result.Titulo)
	})

	t.Run("When isbn not exist must return error", func(t *testing.T) {
		// Arrange
		target := BookServices{}
		target.books = []Book{
			{
				Isbn:   "un isbn",
				Titulo: "Un libro mockeado",
			},
		} // Esta es la implementacion "fake" de los libros y que me sirve solo para este test

		// Act
		_, err := target.SearchByIsbn("otro isbn")

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "Book is not found", err.Error())
	})
}
