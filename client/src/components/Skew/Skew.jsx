import React, { useEffect } from 'react';
import StockChart from '../Chart/StockChart';
import { selectDailyData } from '../../redux/tickers/selectors';
import { getAllDailyData } from '../../redux/tickers/operations';
import { useDispatch, useSelector } from 'react-redux';

const Skew = () => {
    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const dailyData = selectDailyData(selector);

    useEffect(() => {
        dispatch(getAllDailyData('^skew'));
    }, []);

    console.log(dailyData);
    return <StockChart />;
};

export default Skew;
