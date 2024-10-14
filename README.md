# uptime_quardian_bot

## Creating a project

git clone git@github.com:novozhilovsergeydisk/website_checker_bot .

## Initializing the module

go mod init packet-name

## Installing packages

go get github.com/go-telegram-bot-api/telegram-bot-api/v5

## Creating a systemd service

Create a service file in /etc/systemd/system/uptime_quardian_bot.service

Open the file in the editor - sudo nano /etc/systemd/system/uptime_quardian.service

Paste the following code into it

```
[Unit]
Description=Uptime Checker Bot  # Replace with your description
After=network.target

[Service]
Type=simple
User=your_username  # Replace with your username
WorkingDirectory=/path/to/uptime_quard_bot   # Replace with the path to your project
ExecStart=/usr/local/go/bin/go run /path/to/website_checker_bot/cmd/bot/main.go   # Path to your Go binary
Restart=on-failure
Environment="TELEGRAM_BOT_TOKEN=your_telegram_bot_token"  # Set an environment variable with the token

[Install]
WantedBy=multi-user.target
```

#### Explanation of parameters:

>Description — a short description of the service.
After=network.target — the service is started after network services.
User — the user on whose behalf the service will be started.
WorkingDirectory — the working directory for the bot.
ExecStart — the command to start the bot. Make sure the path to your Go binary is correct.
Restart=on-failure — automatically restarts the bot if it fails.
Environment — an environment variable for passing the Telegram token.

## Updating systemd

sudo systemctl daemon-reload

## Autostart service

sudo systemctl enable website_checker_bot.service

sudo systemctl start website_checker_bot.service

## Checking the service status

sudo systemctl status website_checker_bot.service

## Restarting and stopping the service

sudo systemctl restart website_checker_bot.service

sudo systemctl stop website_checker_bot.service

## Logging

sudo journalctl -u website_checker_bot.service