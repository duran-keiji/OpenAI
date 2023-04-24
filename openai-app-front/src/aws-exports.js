const config = {
    Auth: {
        region: process.env.REACT_APP_AWS_AUTH_REGION,
        userPoolId: process.env.REACT_APP_AWS_AUTH_USER_POOL_ID,
        userPoolWebClientId: process.env.REACT_APP_AWS_AUTH_CLIENT_ID
    },
    API: {
        endpoints: [{
            name: "MainApi",
            endpoint: process.env.REACT_APP_OPENAI_APP_API_ENDPOINT,
        }, ]
    }
};

export default config;