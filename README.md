# OpenAI

# ローカル環境の立ち上げ方
①cd /Users/duran_keiji/develop/OpenAI/openai-app-back
②sam build
③sam local start-api
④別のタブを開く
⑤curl "http://127.0.0.1:3000/search/word?q=任意の文字列"

# エラー対応
以下のエラーが出た場合は、
①https://platform.openai.com/account/api-keysにアクセス
②Create new secret keyを押下
③/Users/duran_keiji/develop/OpenAI/openai-app-back/template.yamlのAPI_KEYを最新のものに変更

{
    "error": {
        "message": "Incorrect API key provided: sk-EMmlH***************************************bymc. You can find your API key at https://platform.openai.com/account/api-keys.",
        "type": "invalid_request_error",
        "param": null,
        "code": "invalid_api_key"
    }
}

# Appendix
関数の作り方
sam init --runtime go1.x --name "openai-app-back" --app-template "hello-world"
