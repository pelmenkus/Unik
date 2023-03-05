import telebot

bot = telebot.TeleBot('5856185624:AAG_7V3vO0iBsZdHzbXYN1zfrAb4JvK3Ihg')

@bot.message_handler(commands=['start','help'])
def start(message):
    fn=message.from_user.first_name
    ln=message.from_user.last_name
    if fn==None:
        fn=''
    if ln==None:
        ln=''
    name=f'<b>АААА Негры</b>'+' '+fn+' '+ln
    bot.send_message(message.chat.id,name, parse_mode='html')

@bot.message_handler()
def get_txt(message):
    fn = message.from_user.first_name
    ln = message.from_user.last_name
    if fn == None:
        fn = ''
    if ln == None:
        ln = ''
    answ=fn+' '+ln+' '+message.text+' '+f'<u><b>Пошёл Нахуй</b></u>'
    bot.send_message(message.chat.id, answ, parse_mode='html')

bot.polling(none_stop=True)
