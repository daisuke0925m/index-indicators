import React, { useEffect, useState } from 'react';
import axios from 'axios';
import StockChart from '../Chart/StockChart';

const Skew = () => {
    const [skews, setFgis] = useState([]);

    useEffect(() => {
        async function fetchSkews() {
            try {
                const response = await axios.get('/ticker?symbol=^skew&start=2021-01-10&end=2021-01-25');
                setFgis(response.data);
            } catch (error) {
                console.log(error);
                setFgis([]);
            }
        }
        fetchSkews();
    }, []);
    console.log(skews);

    return (
        <>
            <StockChart />
        </>
    );
};

export default Skew;
