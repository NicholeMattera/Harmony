package store

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StoreLayer interface {
	CreateComposition(createdBy, name string, defaultPermissionFlag PermissionFlag) (Composition, error)
	CreateField(compositionID uint, createdBy, defaultValue string, fieldType FieldType, metadata map[string]interface{}, name string, position uint8, required bool) (Field, error)
	CreateEntity(compositionID uint, createdBy, name string, parentEntityId *uint, slug string) (Entity, error)
	CreateUser(createdBy, email, firstName, lastName, username, password string, roleId uint) (User, error)
	CreateSession(ipAddress string, userId uint) (Session, error)
	CreateRole(createdBy, lockSessionToIp, multipleSessions bool, name string, sessionTimeout *uint64) (Role, error)
	GetAllCompositions() ([]Composition, error)
	GetAllEntities() ([]Entity, error)
	GetAllUsers() ([]User, error)
	GetAllRoles() ([]Role, error)
}

type storeLayer struct {
    db *gorm.DB
}

func New() *storeLayer {
	username, usernameExists := os.LookupEnv("HARMONY_DB_USERNAME")
	if !usernameExists {
		log.Fatal("Error: No database username provided")
		return nil
	}

	password, passwordExists := os.LookupEnv("HARMONY_DB_PASSWORD")
	if !passwordExists {
		log.Fatal("Error: No database password provided")
		return nil
	}

	host, hostExists := os.LookupEnv("HARMONY_DB_HOST")
	if !hostExists {
		log.Fatal("Error: No database hostname provided")
		return nil
	}

	name, nameExists := os.LookupEnv("HARMONY_DB_NAME")
	if !nameExists {
		log.Fatal("Error: No database name provided")
		return nil
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, name)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: Unable to connect to MySQL server", err)
	}

	db.AutoMigrate(&Role{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&Composition{})
	db.AutoMigrate(&Permission{})
	db.AutoMigrate(&Field{})
	db.AutoMigrate(&Entity{})
	db.AutoMigrate(&FieldValue{})

    return &storeLayer{
        db: db,
    }
}
