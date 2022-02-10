# Golang + Dependency Inyection(inverse of control) + TDD

Este es un pequeño ejemplo de una implementacion de TDD + Inyeccion de dependencias ( IOC ) con Golang que he escrito para una publicacion en el blog de [Code4ndreani](https://medium.com/code4ndreani
).


## Instalation
```
git clone https://github.com/victorparedes/golang_di_tdd.git
cd golang_di_tdd.git
go install
go run main.go 0575077220
```

## Run
```
go install

// Must show a existing book
go run main.go 0575077220

// Must show error message
go run main.go 9292929292
```

## Test
```
go test ./...
```

