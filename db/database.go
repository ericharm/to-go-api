package db

import (
    "log"
    "io/ioutil"
    "bytes"
    "fmt"

    "github.com/go-yaml/yaml"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var activeDB *gorm.DB

func InitDB(env string) *gorm.DB {
    gormDb, err := ConnectGorm(env)
    if err != nil {
        panic(err)
    }
    activeDB = gormDb
    return activeDB
}

func GetActiveDB() *gorm.DB {
    return activeDB
}

func ConnectSql(env string) (*sql.DB, error) {
    dbConfig, err := mapConfig()
    driver := dbConfig[env]["driver"]
    db, err := sql.Open(driver, getConnectionString(env, true))
    return db, err
}

func ConnectGorm(env string) (*gorm.DB, error) {
    dbConfig, err := mapConfig()
    driver := dbConfig[env]["driver"]

    db, err := gorm.Open(driver, getConnectionString(env, true))
    return db, err
}

func CreateDatabase(env string) {
    dbConfig, err := mapConfig()
    driver := dbConfig[env]["driver"]

    db, err := sql.Open(driver, getConnectionString(env, false))
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    envConfig := dbConfig[env]
    _, err = db.Exec("CREATE DATABASE " + envConfig["database"])
    if err == nil {
        fmt.Printf("DB '%s' Created.\n", envConfig["database"])
    } else {
        panic(err)
    }

    _, err = db.Exec("USE " + envConfig["database"])
    if err == nil {
        fmt.Printf("Using %s\n", envConfig["database"])
    } else {
        panic(err)
    }
}

func getConnectionString(env string, exists bool) string {
    dbConfig, err := mapConfig();

    if err != nil {
        panic(err)
    }

    envConfig := dbConfig[env]
    list := []string{
        envConfig["user"], ":", envConfig["password"], "@tcp(",
        envConfig["host"], ":", envConfig["port"], ")/",
    }
    if exists {
        list = append(list, envConfig["database"])
    }

    list = append(list, "?parseTime=true")

    var str bytes.Buffer
    for _, l := range list {
        str.WriteString(l)
    }
    return str.String()
}

func mapConfig() (map[string]map[string]string, error) {
    dbYaml, err := ioutil.ReadFile("./config/database.yml")
    dbConfig := make(map[string]map[string]string)
    err = yaml.Unmarshal([]byte(dbYaml), &dbConfig)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return dbConfig, err
}

