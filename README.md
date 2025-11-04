# Books API - Go + SQLite

API REST para gestionar libros y autores con relaciones muchos-a-muchos usando Go y SQLite.

## 🚀 Instalación
```bash
git clone https://github.com/DevJorgeRafael/go-books-sqlite
cd books-sqlite
go mod tidy
```

## ▶️ Ejecutar
```bash
go run main.go
```

Servidor disponible en: `http://localhost:8080`

---

## 📚 Endpoints

### **Books**
- `GET /books` - Listar todos los libros
- `GET /books/:id` - Obtener un libro por ID (incluye autores)
- `POST /books` - Crear un libro
- `PUT /books/:id` - Actualizar un libro
- `DELETE /books/:id` - Eliminar un libro

### **Authors**
- `GET /authors` - Listar todos los autores
- `GET /authors/:id` - Obtener un autor por ID (incluye libros)
- `POST /authors` - Crear un autor
- `PUT /authors/:id` - Actualizar un autor
- `DELETE /authors/:id` - Eliminar un autor

### **Asociaciones (Book-Authors)**
- `POST /author-books` - Asociar un libro con un autor
- `DELETE /author-books?bookId=X&authorId=Y` - Desasociar

---

## 🧪 Ejemplos de uso

### Crear un libro
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "publicationYear": 2008,
    "isbn": "978-0132350884"
  }'
```

### Crear un autor
```bash
curl -X POST http://localhost:8080/authors \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Robert C. Martin",
    "biography": "Software engineer and author",
    "country": "USA"
  }'
```

### Asociar autor a libro
```bash
curl -X POST http://localhost:8080/author-books \
  -H "Content-Type: application/json" \
  -d '{
    "bookId": 1,
    "authorId": 1
  }'
```

### Obtener libro con autores
```bash
curl http://localhost:8080/books/1
```

---

## 📁 Estructura del proyecto
```
books-sqlite/
│
├── main.go                    # Punto de entrada
├── books.db                   # Base de datos SQLite
│
└── internal/
    ├── app/
    │   ├── dependencies.go    # Inyección de dependencias
    │   └── routes.go          # Configuración de rutas
    │
    ├── database/
    │   ├── connection.go      # Conexión a DB
    │   └── migrations.go      # Migraciones SQL
    │
    ├── errors/
    │   └── errors.go          # Errores personalizados
    │
    ├── model/
    │   ├── book.go            # Entidad Book
    │   ├── author.go          # Entidad Author
    │   └── author_book.go     # Tabla intermedia
    │
    ├── service/
    │   ├── book_service.go
    │   ├── author_service.go
    │   └── author_book_service.go
    │
    ├── store/
    │   ├── book_store.go      # Acceso a datos de libros
    │   ├── author_store.go    # Acceso a datos de autores
    │   └── author_book_store.go
    │
    └── transport/
        ├── book_handler.go    # Handlers HTTP de libros
        ├── author_handler.go  # Handlers HTTP de autores
        └── author_book_handler.go
```

---

## 🏗️ Arquitectura

**Layered Architecture** con separación de responsabilidades:

- **Transport** → Capa HTTP (handlers)
- **Service** → Lógica de negocio
- **Store** → Acceso a datos (repository pattern)
- **Model** → Entidades de dominio

---

## 🛠️ Stack tecnológico

- **Go 1.21+**
- **SQLite** (`modernc.org/sqlite`)
- **net/http** (stdlib)

---

## 📝 Notas

Proyecto de aprendizaje enfocado en:
- Clean Architecture básica
- Inyección de dependencias
- Relaciones muchos-a-muchos
- Manejo de errores en Go
- Buenas prácticas REST

---

## 📄 Licencia

MIT