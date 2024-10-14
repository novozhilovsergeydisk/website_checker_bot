package telegram

import (
    "log"
    "time"
    "website_checker_bot/pkg/checker"
    "website_checker_bot/config"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(cfg *config.Config) error {
    bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
    if err != nil {
        return err
    }

    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        if update.Message.IsCommand() {
            switch update.Message.Command() {
            case "check":
                url := update.Message.CommandArguments()
                if url == "" {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide a URL to check.")
                    bot.Send(msg)
                    continue
                }

                status, err := checker.CheckWebsite(url)
                if err != nil {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error checking website: "+err.Error())
                    bot.Send(msg)
                    continue
                }

                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Website "+url+" is "+status)
                bot.Send(msg)
            default:
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command.")
                bot.Send(msg)
            }
        }
    }

    return nil
}

func StartAutomaticCheck(cfg *config.Config) {
    go func() {
        for {
            log.Println("Starting automatic website check...")
            for _, site := range cfg.Websites {
                status, err := checker.CheckWebsite(site)
                if err != nil {
                    log.Printf("Error checking website %s: %v\n", site, err)
                } else {
                    log.Printf("Website %s is %s\n", site, status)
                }
            }
            time.Sleep(cfg.CheckInterval)
        }
    }()
}
