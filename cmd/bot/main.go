package main

import (
    "log"
    "website_checker_bot/internal/telegram"
    "website_checker_bot/config"
)

func main() {
    cfg := config.LoadConfig()

    log.Println("Starting Telegram bot...")

    err := telegram.StartBot(cfg)
    if err != nil {
        log.Fatalf("Error starting Telegram bot: %v", err)
    }

    // Запуск автоматической проверки сайтов
    telegram.StartAutomaticCheck(cfg)
}
