import axios from 'axios';

const httpClient = axios.create({
    baseURL: 'http://localhost:8080',
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
});

httpClient.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        if (error.config && error.response && error.response.status === 401) {
            httpClient
                .post(
                    '/refresh_token',
                    {},
                    {
                        withCredentials: true,
                    }
                )
                .then((res) => {
                    if (res.status == 200) {
                        console.log('success token refresh');
                        return axios.request(error.config);
                    } else {
                        console.log('failure token refresh');
                    }
                })
                .catch((error) => {
                    console.log('failure token refresh', error);
                });
        }
        return Promise.reject(error);
    }
);
export default httpClient;
