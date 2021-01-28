import React, { useEffect, useState } from 'react';
import axios from 'axios';
import StockChart from '../Chart/StockChart';
import { TagSearch } from '../UiKits';

const Comparison = () => {
    const [chartAry, setChartAry] = useState([]);
    const [keyword, setKeyword] = useState([]);
    // const [keywordStore, setKeywordStore] = useState([]);

    const fetchTickers = (symbol) => {
        async function fetchTickers() {
            try {
                const response = await axios.get(`/ticker?symbol=${symbol}`);
                const data = response.data;
                setChartAry([...chartAry, [...data.daily]]);
            } catch (error) {
                setChartAry([]);
            }
        }
        fetchTickers();
    };

    // keywordが追加された時は配列の一番最後のsymbolを返す
    const addingSymbol = (symbols) => {
        return symbols[symbols.length - 1];
    };

    // keywordが削除された時は削除されたkeywordのindexを探してChartAryから該当する配列を削除する
    const reduceTicker = () => {
        for (let i = 0; i < chartAry.length; i++) {
            const reducedKW = keyword[i];
            const reducingSymbol = chartAry[i][0].symbol;
            if (reducedKW !== reducingSymbol) {
                const newChartAry = [...chartAry];
                newChartAry.splice(i - 1, 1);
                setChartAry(newChartAry);
            }
        }
    };

    useEffect(() => {
        if (keyword.length > chartAry.length) {
            fetchTickers(addingSymbol(keyword));
        } else {
            reduceTicker();
        }
    }, [keyword]);

    return (
        <section>
            <TagSearch setKeyword={setKeyword} />
            {chartAry.length ? <StockChart chartAry={chartAry} title={'Compare Chart '} /> : 'loading'}
        </section>
    );
};

export default Comparison;
