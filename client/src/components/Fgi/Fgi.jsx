import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Fgi = () => {
    const [fgi, setFgis] = useState([]);

    useEffect(() => {
        async function fetchFgis() {
            try {
                const response = await axios.get('/fgi');
                setFgis(response.data);
            } catch (error) {
                console.log(error);
                setFgis([]);
            }
        }
        fetchFgis();
    }, []);

    return <div>{fgi.length ? fgi.map((d, i) => <div key={i}>{d.now_value}</div>) : 'loading'}</div>;
};

export default Fgi;
