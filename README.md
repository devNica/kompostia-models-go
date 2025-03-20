# kompostia-models-go

# 📚 Características

    Modelos definidos utilizando GORM.

    Inicialización sencilla de la conexión a PostgreSQL.

    Migraciones automáticas de los modelos.

    Modularidad para utilizar el paquete en otros proyectos de Go.

# 📂 Estructura del Proyecto

    kompostia-models-go/
    ├── go.mod
    ├── go.sum
    ├── pkg/
    |   ├── schemas/
    |   |    ├── brand.go
    |   |    ├── category.go
    |   |    ├── db.go
    |   |    ├── file.go
    |   |    ├── item.go
    |   |    ├── location.go
    |   |    ├── settings.go
    |   |    └── supplier.go   
    │   ├── models/
    |   |    ├── queries/
    |   |    |    ├── hierarchical_category_relationship.sql
    |   |    |    └── registered_categories.sql
    |   |    ├── query_loader.go
    |   |    ├── sql_queries.go
    |   |    ├── brand.go
    |   |    ├── category.go
    |   |    ├── file.go
    |   |    ├── item.go
    |   |    ├── location.go
    │   │    └── supplier.go
    |   ├── sahred/
    |   |    └── generics.go
    │   └── conn/
    │        └── database.conn.go
    ├── LICENSE
    └── README.md


## 🚀 Instalación

Asegúrate de tener configurado Go y ejecuta el siguiente comando en tu proyecto:

`go get github.com/devnica/kompostia-models-go`

# 📌 Uso

## 1. Importar el Paquete

`import ("github.com/devnica/kompostia-models-go/pkg/conn")`


## 2. Inicializar la Conexión a la Base de Datos

    package main

    import (
        "fmt"

        schema "github.com/devnica/kompostia-models-go/pkg/schemas"
        "github.com/devnica/kompostia-models-go/pkg/conn"
    )

    func main() {

        dbParams := schema.DbSchema{
            DatabaseName:     "dbName",
            DatabasePassword: "dbPassword",
            DatabaseUser:     "dbUser",
            DatabaseHost:     "dbHost",
            DatabasePort:     "dbPort",
            DatabaseSchema:   "dbSchema",
            DatabaseSSLMode:  "dbSSL",
        }

        var refreshModels = true

        _, err := conn.InitDB(dbParams, refreshModels)

        if err != nil {
            panic(fmt.Sprintf("Error: %v", err))
        }

    }

# 📖 Funciones Exportadas

    func InitDB(params schema.DbSchema, refreshModels bool) (*gorm.DB, error)

    Descripción: Inicia la conexión a la base de datos y realiza las migraciones.

    Parámetros:

    params ({
        DatabaseName     string
        DatabasePassword string
        DatabaseUser     string
        DatabaseHost     string
        DatabasePort     string
        DatabaseSchema   string
        DatabaseSSLMode  string
    }): parametros de conexión a PostgreSQL.

    Retorna:

    *gorm.DB: Instancia de la conexión a la base de datos.

    error: Error si la conexión falla.

    func migrateModels(db *gorm.DB, model interface{}, modelName string) error

    Descripción: Ejecuta la migración de un modelo específico.

    Parámetros:

    db (*gorm.DB): Conexión activa a la base de datos.

    model (interface{}): Estructura del modelo.

    modelName (string): Nombre del modelo para mensajes de logs.

    Retorna:

    error: Error si la migración falla.

# 📜 Licencia

Este proyecto está bajo la licencia Apache 2.0.
