import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Fgi = () => {
    const [fgi, setFgi] = useState([]);

    useEffect(() => {
        async function fetchFgi() {
            try {
                const response = await axios.get('/fgi');
                setFgi(response.fgi);
            } catch (error) {
                console.log(error);
                setFgi([]);
            }
        }
        fetchFgi();
    }, []);

    return <div>{fgi.length ? fgi.map((d, i) => <div key={i}>{d.now_value}</div>) : 'loading'}</div>;
};

export default Fgi;
