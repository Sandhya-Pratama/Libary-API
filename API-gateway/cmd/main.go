package main

import (
    "log"
    "net/http"


)

func main() {
    cfg := config.LoadConfig()

    r := router.SetupRouter(cfg)
    
    log.Println("Starting API Gateway on port", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Failed to start API Gateway: %v", err)
    }
}
