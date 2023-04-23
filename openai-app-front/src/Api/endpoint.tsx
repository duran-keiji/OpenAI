export const BASE_ENDPOINT = process.env.REACT_APP_OPENAI_APP_API_ENDPOINT;

// #################### API一覧 ####################

export const API_ENDPOINTS = {
  chatgptSearchWord: "/search/word",
//   getUserById: "users/:id",

};

// #################################################

export const getApi = (endpoint: string) => `${BASE_ENDPOINT}${endpoint}`;
