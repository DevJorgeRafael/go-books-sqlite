# Books API - Go + SQLite

API REST simple para gestionar libros usando Go y SQLite.

## Instalación
```bash
git clone <tu-repo>
cd books-sqlite
go mod tidy
```

## Ejecutar
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## Endpoints

- `GET /books` - Obtener todos los libros
- `GET /books/:id` - Obtener un libro por ID
- `POST /books` - Crear un libro
- `PUT /books/:id` - Actualizar un libro
- `DELETE /books/:id` - Eliminar un libro

## Ejemplo de uso
```bash
# Crear un libro
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"titulo":"1984","autor":"George Orwell"}'

# Obtener todos los libros
curl http://localhost:8080/books
```

## Estructura

- `model/` - Estructuras de datos
- `store/` - Capa de acceso a datos
- `service/` - Lógica de negocio
- `transport/` - Handlers HTTP