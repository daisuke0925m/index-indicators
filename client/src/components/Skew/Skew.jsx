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

    return (
        <section>{dailyData.daily.length ? <StockChart daily={dailyData.daily} title={'SKEW'} /> : 'loading'}</section>
    );
};

export default Skew;
