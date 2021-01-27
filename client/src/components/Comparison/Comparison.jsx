import React, { useEffect, useState } from 'react';
import StockChart from '../Chart/StockChart';
import { selectDailyData } from '../../redux/tickers/selectors';
import { getAllDailyData } from '../../redux/tickers/operations';
import { useDispatch, useSelector } from 'react-redux';

const Comparison = () => {
    const [chartAry, setChartAry] = useState([[{ symbol: '', date: '', close: 0 }]]);
    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const dailyData = selectDailyData(selector);

    useEffect(() => {
        dispatch(getAllDailyData('spy'));
        setChartAry([[...dailyData.daily]]);
    }, []);

    const addTicker = () => {
        dispatch(getAllDailyData('spxl'));
        setChartAry([...chartAry, [...dailyData.daily]]);
    };

    const reduceTicker = () => {
        const len = chartAry.length;
        chartAry.splice(len - 1, 1);
        setChartAry(chartAry);
    };

    console.log(chartAry);
    console.log(chartAry.length);

    return (
        <section>
            <button onClick={() => addTicker()}>追加ボタン</button>
            <button onClick={() => reduceTicker()}>削除ボタン</button>
            {chartAry.length ? <StockChart chartAry={chartAry} title={'Compare Chart '} /> : 'loading'}
        </section>
    );
};

export default Comparison;
