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

func (db *DBProxy) GetBookTrackerUser() *models.User {
    user := models.User{}
    db.db.Where("name = ?", "booktracker").First(&user)
    return &user
}

func (dbp *DBProxy) migrate() {
    dbp.db.AutoMigrate(&models.Book{})
    dbp.db.AutoMigrate(&models.Reading{})
    dbp.db.AutoMigrate(&models.Session{})
    dbp.db.AutoMigrate(&models.Shelf{})
    dbp.db.AutoMigrate(&models.User{})

    dbp.createSampleData()
}

func (dbp *DBProxy) clearSessions() {
   dbp.db.Exec("DELETE FROM sessions;")
}

// create a 'booktracker' user and a read, reading, to-read shelves for sample data in the app
// this will assist for adding books and viewing lists until account creation is fully implemented
func (dbp *DBProxy) createSampleData() {
    user := models.User{
        Login: "booktracker",
        Email: "booktracker@gmail.com",
        Password: "booktracker",
    }
    dbp.db.FirstOrCreate(&user, user)

    shelves := [3]string{"Read", "Reading", "ToRead"}
    for _, shelfName := range shelves {
        shelf := models.Shelf {
            Name: shelfName,
            UserId: user.ID,
        }
        dbp.db.FirstOrCreate(&shelf, shelf)
    }
}
