import React, { useEffect, useState } from 'react';
import httpClient from '../../axios';
import StockChart from '../Chart/StockChart';

const Comparison = () => {
    const [chartAry, setChartAry] = useState([]);

    useEffect(() => {
        async function fetchTickers() {
            try {
                const response = await httpClient.get(`/ticker?symbol=^skew`);
                const data = response.data;
                setChartAry([...chartAry, [...data.daily]]);
            } catch (error) {
                console.log(error);
                setChartAry([]);
            }
        }
        fetchTickers();
    }, []);

    return (
        <section>
            {chartAry.length ? <StockChart chartAry={chartAry} title={'Skew'} chartColor={'blue'} /> : 'loading'}
        </section>
    );
};

export default Comparison;
