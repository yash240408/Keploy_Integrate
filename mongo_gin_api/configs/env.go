package configs

// What this file does is :-
// Import the required dependencies.
// Create an EnvMongoURI function that checks if the environment variable is correctly loaded and returns the environment variable.

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}