package main

import (
    "encoding/json"
    "os"
)

func startWriter(results chan check, outputPath string) {
    file, err := os.Create(outputPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)

    for result := range results {
        encoder.Encode(result)
    }
}