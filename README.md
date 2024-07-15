# Raybot

Telegram bot on Raydium written in Go

## Features

- Get all liquidity pools on Raydium
- Get all concentrated liquidity pools on Raydium
- Get all standard liquidity pools on Raydium

## Project structure

This project follows structure of [Golang standard](https://github.com/golang-standards/project-layout). The project is structured as follows:

- `cmd`: contains the main package
- `internal`: contains the main logic of the bot
    - `app`: your actual application code
    - `conf`: load configuration environment
    - `handler`: view in MVC pattern. It is responsible for handling the incoming bot requests
    - `entity`: model in MVC pattern
    - `service`: controller in MVC pattern. It is responsible for handling the business logic, such as get all pools on Raydium from api.

## Installation

1. Get your bot token from BotFather
2. Create a `.env` file in the root directory and add the following line:
```
BOT_TOKEN=<your_bot_token>
```
3. Run the bot with docker
```bash
docker compose up -d
```

