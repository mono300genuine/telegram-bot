import os
import telebot
from pycoingecko import CoinGeckoAPI

BOT_TOKEN = os.environ.get('BOT_TOKEN')

bot = telebot.TeleBot(BOT_TOKEN)
cg = CoinGeckoAPI()

@bot.message_handler(commands=['start', 'hello'])
def send_welcome(message):
    bot.reply_to(message, "Howdy, how are you doing?")

@bot.message_handler(commands=['recent'])
def recent_tokens_info(message):
    try:
        # Fetch the 10 most recently added tokens
        recent_tokens = cg.get_coins_list()
        recent_tokens = sorted(recent_tokens, key=lambda x: x['id'], reverse=True)[:10]
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
                contact_info = "🚫 No contact information available."
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
        
        bot.reply_to(message, response)
    except Exception as e:
        bot.reply_to(message, f"Error retrieving data: {e}")

@bot.message_handler(func=lambda msg: True)
def echo_all(message):
    bot.reply_to(message, message.text)
    print("bot message:", message.text)

bot.infinity_polling()
