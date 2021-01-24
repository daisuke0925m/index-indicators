import React, { useEffect, useState } from 'react';
import axios from 'axios';
import SimpleChart from '../UiKits/SimpleChart';

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

    return <div>{fgi.length ? <SimpleChart /> : 'loading'}</div>;
};

export default Fgi;
