import React, { useEffect, useState } from 'react'
import { API, Auth } from 'aws-amplify'

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

        const apiName = 'MainApi';
        const path = '/search/word' + query;
        const myInit = {
          headers: { 
            Authorization: `Bearer ${(await Auth.currentSession()).getIdToken().getJwtToken()}`,
          },
        };
      
        const resultText = await API.get(apiName, path, myInit);
        setResultText(resultText)
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
