import os
import time
import telebot
import matplotlib.pyplot as plt
from pycoingecko import CoinGeckoAPI
from datetime import datetime

BOT_TOKEN = os.environ.get('BOT_TOKEN')

bot = telebot.TeleBot(BOT_TOKEN)
cg = CoinGeckoAPI()

@bot.message_handler(commands=['start', 'hello'])
def send_welcome(message):
    bot.reply_to(message, "👋 Hello, how are you doing?")

@bot.message_handler(commands=['recent'])
def recent_tokens_info(message):
    start_time = time.time()
    
    try:
        # Fetch the 10 most recently added tokens
        recent_tokens = cg.get_coins_list()
        recent_tokens = sorted(recent_tokens, key=lambda x: x['id'], reverse=True)[:3]
        token_ids = [token['id'] for token in recent_tokens]

        # Fetch current prices and 24-hour trading volumes for these tokens
        token_data = cg.get_price(ids=','.join(token_ids), vs_currencies='usd', include_24hr_vol=True)
        
        # Fetch detailed information for the tokens to get the official website and community links
        token_details = {token['id']: cg.get_coin_by_id(token['id']) for token in recent_tokens}

        response = "Recent Tokens:\n"
        for token in recent_tokens:
            token_id = token['id']
            token_name = token['name']
            token_symbol = token['symbol']
            price = token_data[token_id]['usd'] if token_id in token_data else 'N/A'
            volume = token_data[token_id]['usd_24h_vol'] if token_id in token_data else 'N/A'
            details = token_details.get(token_id, {})
            website = details.get('links', {}).get('homepage', ['N/A'])[0]
            telegram = details.get('links', {}).get('telegram_channel_identifier', 'N/A')
            discord = details.get('links', {}).get('chat_url', ['N/A'])[0]

            if website == 'N/A' and telegram == 'N/A' and discord == 'N/A':
                contact_info = "No contact information available."
            else:
                contact_info = (
                    f"  Website: {website}\n"
                    f"  Telegram: {telegram}\n"
                    f"  Discord: {discord}\n"
                )

            response += (
                f"{token_name} ({token_symbol}):\n"
                f"  Price: ${price}\n"
                f"  24h Volume: ${volume}\n"
                f"{contact_info}\n\n"
            )

        # Measure response time
        response_time = time.time() - start_time
        response += f"Response Time: {response_time:.2f} seconds\n"

        bot.reply_to(message, response)

        # Generate and send price chart for each token
        for token in recent_tokens:
            token_id = token['id']
            token_name = token['name']
            token_symbol = token['symbol']

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

    except Exception as e:
        bot.reply_to(message, f"Error retrieving data: {e}")

@bot.message_handler(func=lambda msg: True)
def echo_all(message):
    bot.reply_to(message, message.text)
    print("bot message:", message.text)

bot.infinity_polling()
