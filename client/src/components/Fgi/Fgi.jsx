import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Fgi = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        async function fetchData() {
            try {
                const response = await axios.get('/fgi');
                setData(response.data);
            } catch (error) {
                console.log(error);
                setData([]);
            }
        }
        fetchData();
    }, []);

    return <div>{data.length ? data.map((d, i) => <div key={i}>{d.now_value}</div>) : 'loading'}</div>;
};

export default Fgi;
