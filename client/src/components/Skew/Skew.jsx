import React, { useEffect, useState } from 'react';
import axios from 'axios';
import StockChart from '../Chart/StockChart';
import { tickerDateParse } from '../Functions/functions';

const Skew = () => {
    const [skews, setFgis] = useState([]);
    const today = new Date();
    const parsedToday = tickerDateParse(today, 0);
    const parsed1mAgo = tickerDateParse(today, -50);

    useEffect(() => {
        async function fetchSkews() {
            try {
                const response = await axios.get(`/ticker?symbol=^skew&start=${parsed1mAgo}&end=${parsedToday}`);
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
