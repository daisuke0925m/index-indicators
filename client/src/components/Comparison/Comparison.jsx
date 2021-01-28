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
    // keywordが削除された時は削除されたkeywordのindexを探してChartAryから該当する配列を削除する
    const addSymbols = (symbols) => {
        return symbols[symbols.length - 1];
    };

    useEffect(() => {
        if (keyword.length > chartAry.length) {
            fetchTickers(addSymbols(keyword));
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
