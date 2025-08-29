package env

import (
    "log"
    "os"
    "strconv"
)

func GetStringEnv(key string) string {
    value := GetEnv(key)
    if value == "" {
        log.Printf("Env variable unable to load")
        os.Exit(1)
    }

    return value
}

func GetIntEnv(key string) int {
    value := GetEnv(key)
    if value == "" {
        log.Printf("Env variable unable to load")
        os.Exit(1)
    }

    intValue , err := strconv.Atoi(value)
    if err != nil {
        log.Printf("unable to parse string to int ")
        os.Exit(1)
    }
    return intValue
    
}
