# biblioteca
API REST de un CRUD de libros utilizando patrón MVC en Go y PostgreSQL

1. En la carpeta "files", se encuentra el script "librarydb_postgresql.sql" para crear la base de datos.
2. En el archivo "model/connectDB.go", actualizar los datos para la conexión a la base de datos.
3. Abrir una terminal y ubicarse en la raiz del proyecto, por ejemplo, cd ~/[ruta proyecto] 
4. Ejecutar el comando go run .
5. Desde un cliente API REST, por ejemplo, Postman, importar el archivo "BIBLIOTECA.postman_collection.json"
6. En la raiz del proyecto crear el archivo: .env, adicionar las variables de entorno con los valores correspondientes a tu entorno de desarrollo
    DB_USER=root
    DB_PASS=
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=librarydb