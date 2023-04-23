import React, { useEffect, useState } from 'react'
// import { API } from 'aws-amplify'
import { API_ENDPOINTS, getApi } from '../Api/endpoint';


const Home: React.FC = () => {
  const [text, setText] = useState<string>("");
  const [resultText, setResultText] = useState<string>("");


  const handleTextChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value;
    setText(text)
  };


  const chatgptSearchWord = async () => {
    try {
      const query = "?q=" + text
      const requestUrl = getApi(API_ENDPOINTS.chatgptSearchWord) + query
      const response = await fetch(requestUrl);
      const resultText = await response.json();
      setResultText(resultText)
      console.log(resultText);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="App">
      <p>ようこそ！</p>
      <p>chatGPT</p>
      <div>
        <input type="text" value={text} onChange={handleTextChange} />
        <button onClick={chatgptSearchWord}>質問</button>
        <p>{resultText}</p>
      </div>
    </div>
  );
}

export default Home;
