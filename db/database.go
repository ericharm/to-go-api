package db

import (
    "log"
    "io/ioutil"
    "bytes"

    "github.com/go-yaml/yaml"
)

func GetConnectionString(env string, exists bool) string {
    dbConfig, err := MapConfig();

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

    var str bytes.Buffer
    for _, l := range list {
        str.WriteString(l)
    }
    return str.String()
}

func MapConfig() (map[string]map[string]string, error) {
    dbYaml, err := ioutil.ReadFile("./config/database.yml")
    dbConfig := make(map[string]map[string]string)
    err = yaml.Unmarshal([]byte(dbYaml), &dbConfig)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return dbConfig, err
}

