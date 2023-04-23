import React, { useEffect, useState } from 'react'
import { API } from 'aws-amplify'

type SampleDataType = [
  {
    id: string;
    name: string;
  }
];
const initialSampleData = {id: "", name: ""};

const Home: React.FC = () => {
  const [sampleData, setSampleData] = useState<SampleDataType>([initialSampleData])

  useEffect(() => {
    getSampleData()
  }, [])

  const getSampleData = async () => {
    try {
      const API_NAME = 'OpenaiAppApi';
      const PATH = '/search/word?q=APIとは';
      const myInit = {};
    
      const data = await API.get(API_NAME, PATH, myInit);
      setSampleData(data)
    } catch (err) { console.log('error getting data') }
  }

  return (
    <div className="App">
      <p>ようこそ！</p>
      <p>登録データ一覧</p>
      {
        sampleData.map((data, index) => (
          <div key={data.id ? data.id : index}>
            <p>{data.name}</p>
          </div>
        ))
      }
    </div>
  );
}

export default Home;
