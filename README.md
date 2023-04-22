# OpenAI

ローカル環境の立ち上げ方
①cd /Users/duran_keiji/develop/OpenAI/lambda-chatgpt
②sam build
③sam local start-api
④別のタブを開く
⑤curl "http://127.0.0.1:3000/search/word?q=任意の文字列"