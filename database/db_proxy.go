package database

import (
    "github.com/jinzhu/gorm"
    "github.com/mortonar/acs560_course_project/database/models"
)

type DBProxy struct {
  db *gorm.DB
}

func NewDBProxy() *DBProxy {
    args := "user=booktracker dbname=booktracker sslmode=disable password=booktracker"
    db, err := gorm.Open("postgres", args)
    if err != nil {
        return nil // TODO do real error-handling
    }
    dbp := &DBProxy{db}
    dbp.migrate()
    dbp.clearSessions()
    return dbp
}

func (db *DBProxy) GetConnection() *gorm.DB {
    return db.db
}

func (dbp *DBProxy) migrate() {
    dbp.db.AutoMigrate(&models.Book{})
    dbp.db.AutoMigrate(&models.Reading{})
    dbp.db.AutoMigrate(&models.Session{})
    dbp.db.AutoMigrate(&models.Shelf{})
    dbp.db.AutoMigrate(&models.User{})
}

func (dbp *DBProxy) clearSessions() {
   dbp.db.Exec("DELETE FROM sessions;")
}
