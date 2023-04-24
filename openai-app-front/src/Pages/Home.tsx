import React, { useEffect, useState } from 'react'
import { API, Auth } from 'aws-amplify'
import { CognitoUserAmplify } from '@aws-amplify/ui';
import '@aws-amplify/ui-react/styles.css';

type HomeProps = {
  signOut: VoidFunction;
  user: CognitoUserAmplify;
};

const Home: React.FC<HomeProps> = (props) => {
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
      <p>ようこそ、 {props.user.username} さん！</p>
      <p>chatGPT</p>
      <div>
        <input type="text" value={text} onChange={handleTextChange} />
        <button onClick={chatgptSearchWord}>質問</button>
        <p>{resultText}</p>
        <button onClick={props.signOut}>Sign out</button>
      </div>
    </div>
  );
}

export default Home;
