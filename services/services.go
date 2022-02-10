package services

/*
Esta interface existe para que, si necesitaramos, pudieramos mockear este servicio.
La forma de mockearlo es crear una structure que cumpla con la interface y usarlo
en nuestros test segun necesitemos.
*/
type Services interface {
	SearchByIsbn(isbn string) (*Book, error)
}

/*
Este es el DTO/modelo de mis libros
lo deje aca... ¿Por que?... no hay porque...
*/
type Book struct {
	Isbn   string
	Titulo string
}
