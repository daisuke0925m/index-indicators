import React, { useEffect, useState } from 'react';
import axios from 'axios';
import StockChart from '../Chart/StockChart';
import { TagSearch } from '../UiKits';

const Comparison = () => {
    const [chartAry, setChartAry] = useState([]);

    const fetchTickers = (symbol) => {
        async function fetchTickers() {
            try {
                const response = await axios.get(`/ticker?symbol=${symbol}`);
                const data = response.data;
                setChartAry([...chartAry, [...data.daily]]);
            } catch (error) {
                console.log(error);
                setChartAry([]);
            }
        }
        fetchTickers();
    };

    useEffect(() => {
        fetchTickers('spy');
    }, []);

    const addTicker = () => {
        fetchTickers('tlt');
    };

    const reduceTicker = () => {
        const newChartAry = [...chartAry];
        const len = newChartAry.length;
        newChartAry.splice(len - 1, 1);
        setChartAry(newChartAry);
    };

    console.log(chartAry);
    console.log(chartAry.length);

    return (
        <section>
            <TagSearch />
            <button onClick={() => addTicker()}>追加ボタン</button>
            <button onClick={() => reduceTicker()}>削除ボタン</button>
            {chartAry.length ? <StockChart chartAry={chartAry} title={'Compare Chart '} /> : 'loading'}
        </section>
    );
};

export default Comparison;
