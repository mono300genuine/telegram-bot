import os
import time
import asyncio
import aiohttp
import telebot
import matplotlib.pyplot as plt
from pycoingecko import CoinGeckoAPI
from datetime import datetime

BOT_TOKEN = os.environ.get('BOT_TOKEN')

bot = telebot.TeleBot(BOT_TOKEN)
cg = CoinGeckoAPI()

async def fetch_url(session, url):
    async with session.get(url) as response:
        return await response.json()

async def fetch_data():
    async with aiohttp.ClientSession() as session:
        sol_data = await fetch_url(session, 'https://api.coingecko.com/api/v3/coins/solana')
        return sol_data

@bot.message_handler(commands=['start', 'hello'])
def send_welcome(message):
    bot.reply_to(message, "Howdy, how are you doing?")

@bot.message_handler(commands=['sol'])
def sol_token_info(message):
    start_time = time.time()

    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)
    sol_data = loop.run_until_complete(fetch_data())

    token_id = sol_data['id']
    token_name = sol_data['name']
    token_symbol = sol_data['symbol']
    price = sol_data['market_data']['current_price']['usd']
    volume = sol_data['market_data']['total_volume']['usd']
    website = sol_data['links']['homepage'][0]
    telegram = sol_data['links']['telegram_channel_identifier']
    discord = sol_data['links']['chat_url'][0]

    response = (
        f"{token_name} ({token_symbol}):\n"
        f"  Price: ${price}\n"
        f"  24h Volume: ${volume}\n"
        f"  Website: {website}\n"
        f"  Telegram: {telegram}\n"
        f"  Discord: {discord}\n"
    )

    # Measure response time
    response_time = time.time() - start_time
    response += f"Response Time: {response_time:.2f} seconds\n"

    bot.reply_to(message, response)

    try:
        # Fetch historical price data (e.g., last 30 days)
        historical_data = cg.get_coin_market_chart_by_id(id=token_id, vs_currency='usd', days=30)
        prices = historical_data['prices']

        # Extract dates and prices
        dates = [datetime.fromtimestamp(price[0] / 1000) for price in prices]
        values = [price[1] for price in prices]

        # Plot the price chart
        plt.figure(figsize=(10, 5))
        plt.plot(dates, values, label=f'{token_name} (USD)')
        plt.title(f'{token_name} Price Chart (Last 30 Days)')
        plt.xlabel('Date')
        plt.ylabel('Price (USD)')
        plt.legend()
        plt.grid(True)

        # Save the chart as an image
        chart_filename = f'{token_id}_chart.png'
        plt.savefig(chart_filename)
        plt.close()

        # Send the chart image to the user
        with open(chart_filename, 'rb') as chart_file:
            bot.send_photo(message.chat.id, chart_file, caption=f'{token_name} (Last 30 Days)')

    except Exception as e:
        bot.reply_to(message, f"Error generating chart for {token_name}: {e}")

@bot.message_handler(commands=['buy'])
def buy_token(message):
    try:
        # Extract command arguments (quantity)
        args = message.text.split()
        if len(args) != 2:
            bot.reply_to(message, "Usage: /buy <quantity>")
            return

        quantity = float(args[1])

        # Mock buy order
        order = {
            "symbol": "SOL",
            "quantity": quantity,
            "price": 100.0,  # Mock price
            "status": "FILLED",
            "type": "MARKET",
            "side": "BUY"
        }
        bot.reply_to(message, f"Mock Buy Order Placed: {order}")

    except Exception as e:
        bot.reply_to(message, f"Error placing buy order: {e}")

@bot.message_handler(commands=['sell'])
def sell_token(message):
    try:
        # Extract command arguments (quantity)
        args = message.text.split()
        if len(args) != 2:
            bot.reply_to(message, "Usage: /sell <quantity>")
            return

        quantity = float(args[1])

        # Mock sell order
        order = {
            "symbol": "SOL",
            "quantity": quantity,
            "price": 100.0,  # Mock price
            "status": "FILLED",
            "type": "MARKET",
            "side": "SELL"
        }
        bot.reply_to(message, f"Mock Sell Order Placed: {order}")

    except Exception as e:
        bot.reply_to(message, f"Error placing sell order: {e}")

@bot.message_handler(commands=['order_status'])
def order_status(message):
    try:
        # Extract command arguments (order_id)
        args = message.text.split()
        if len(args) != 2:
            bot.reply_to(message, "Usage: /order_status <order_id>")
            return

        order_id = args[1]

        # Mock order status
        order = {
            "orderId": order_id,
            "symbol": "SOL",
            "status": "FILLED",
            "price": 100.0,  # Mock price
            "origQty": 1.0,  # Mock quantity
            "executedQty": 1.0,
            "type": "MARKET",
            "side": "BUY"
        }
        bot.reply_to(message, f"Mock Order Status: {order}")

    except Exception as e:
        bot.reply_to(message, f"Error checking order status: {e}")

@bot.message_handler(func=lambda msg: True)
def echo_all(message):
    bot.reply_to(message, message.text)
    print("bot message:", message.text)

bot.infinity_polling()
