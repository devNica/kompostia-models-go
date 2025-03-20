package conn

import (
	"fmt"

	"github.com/devnica/kompostia-models-go/pkg/models"
	sch "github.com/devnica/kompostia-models-go/pkg/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/* Funcion de Inicializacion de Conexion a la base de Datos*/
func InitDB(params sch.DbSchema, refreshModels bool) (*gorm.DB, error) {

	// leer parametros de conexion a la base de datos
	db_user := params.DatabaseUser
	db_password := params.DatabasePassword
	db_host := params.DatabaseHost
	db_port := params.DatabasePort
	db_name := params.DatabaseName
	ssl_mode := params.DatabaseSSLMode

	db_schema := params.DatabaseSchema

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s", db_host, db_port, db_user, db_password, db_name, ssl_mode, db_schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Fallo la conexion a la base de datos: %w", err)
	}

	var schemaExists bool

	// Verificar que el Schema Existe
	err = db.Raw("SELECT EXISTS(SELECT 1 FROM pg_namespace WHERE nspname = ?)", db_schema).Scan(&schemaExists).Error
	if err != nil {
		return nil, fmt.Errorf("Error al verificar la existencia del schema: %w", err)
	}

	if !schemaExists {
		// Si el schema no existe, crearlo y realizar la migracion de los modelos definidos
		fmt.Printf("El schema '%s' no existe. Creando el schema y realizando migraciones de modelos...\n", db_schema)

		err = db.Exec(fmt.Sprintf("CREATE SCHEMA %s", db_schema)).Error

		if err != nil {
			return nil, fmt.Errorf("Error al crear el schema: %w", err)
		}

		migrateModels(db, &models.SupplierModel{}, "Supplier Model")
		migrateModels(db, &models.ItemBrandModel{}, "ItemBrand Model")
		migrateModels(db, &models.CategoryModel{}, "Category Model")

		migrateModels(db, &models.LocationTypeModel{}, "LocationType Model")
		migrateModels(db, &models.StorageLocationModel{}, "StorageLocation Model")

		migrateModels(db, &models.FileModel{}, "File Model")

		migrateModels(db, &models.CatalogItemModel{}, "CtgItem Model")
		migrateModels(db, &models.ItemHasFileModel{}, "ItemHasFile Model")
		migrateModels(db, &models.ItemHasLocationModel{}, "ItemHasLocation Model")
	}

	// Refrescar los modelos de la base de datos
	if refreshModels {

		migrateModels(db, &models.SupplierModel{}, "Supplier Model")
		migrateModels(db, &models.ItemBrandModel{}, "ItemBrand Model")
		migrateModels(db, &models.CategoryModel{}, "Category Model")

		migrateModels(db, &models.LocationTypeModel{}, "LocationType Model")
		migrateModels(db, &models.StorageLocationModel{}, "StorageLocation Model")

		migrateModels(db, &models.FileModel{}, "File Model")

		migrateModels(db, &models.CatalogItemModel{}, "CtgItem Model")
		migrateModels(db, &models.ItemHasFileModel{}, "ItemHasFile Model")
		migrateModels(db, &models.ItemHasLocationModel{}, "ItemHasLocation Model")
	}

	fmt.Println("Conexion a la base de datos exitosa")

	return db, nil

}

func migrateModels(db *gorm.DB, model interface{}, modelName string) error {

	if err := db.AutoMigrate(model); err != nil {
		fmt.Printf("Error al migrar %s: %v\n", modelName, err)
		return err
	}
	fmt.Printf("Migracion de %s realizada \n", modelName)
	return nil
}
