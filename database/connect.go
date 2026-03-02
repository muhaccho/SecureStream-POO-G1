package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func ObtenerConexion() *sql.DB {
    // Estructura: "usuario:contraseña@tcp(host:puerto)/nombre_bd"
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/securestream_db")
    if err != nil {
        log.Fatal("Error al conectar: ", err)
    }
    return db
}