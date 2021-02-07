import { Button } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import httpClient from '../../axios';
import { getUsersLikes } from '../../redux/users/selectors';
import StockChart from '../Chart/StockChart';
import { SpaceRow, TagSearch } from '../UiKits';

const Comparison = () => {
    const [chartAry, setChartAry] = useState([]);
    const [keywords, setKeywords] = useState([]);
    const [isRegisterBtn, setIsRegisterBtn] = useState(false);
    const selector = useSelector((state) => state);
    const likes = getUsersLikes(selector);

    const fetchTickers = (symbol) => {
        async function fetch() {
            try {
                const response = await httpClient.get(`/ticker?symbol=${symbol}`);
                const data = response.data;
                setChartAry([...chartAry, [...data.daily]]);
            } catch (error) {
                setChartAry([]);
            }
        }
        fetch();
    };

    // keywordsが追加された時は配列の一番最後のsymbolを返す
    const addingSymbol = (symbols) => {
        return symbols[symbols.length - 1];
    };

    // keywordsが削除された時は削除されたkeywordsのindexを探してChartAryから該当する配列を削除する
    const reduceTicker = () => {
        for (let i = 0; i < chartAry.length; i++) {
            const reducedKW = keywords[i];
            const reducingSymbol = chartAry[i][0].symbol;
            if (reducedKW !== reducingSymbol) {
                const newChartAry = [...chartAry];
                newChartAry.splice(i - 1, 1);
                setChartAry(newChartAry);
            }
        }
    };
    // console.log("chartAry", chartAry)
    // console.log("keywords", keywords)
    // console.log("likes", likes)
    // console.log("isRegisterBtn", isRegisterBtn)

    const setRegisteredTickers = () => {
        const registered = likes.reduce((newAry, like) => {
            if (like.symbol !== 'fgi') {
                newAry.push(like.symbol);
            }
            return newAry;
        }, []);
        setIsRegisterBtn(true);
        setChartAry([]);
        setKeywords(registered);
    };

    const setRegisteredChartAry = () => {
        const ary = [];
        async function fetch() {
            try {
                for (let i = 0; i < keywords.length; i++) {
                    const symbol = keywords[i];
                    const response = await httpClient.get(`/ticker?symbol=${symbol}`);
                    const data = response.data;
                    ary.push(data.daily);
                }
                setChartAry(ary);
                return;
            } catch (error) {
                console.error(error);
            }
        }
        fetch();
    };

    useEffect(() => {
        if (keywords.length > chartAry.length && !isRegisterBtn) {
            fetchTickers(addingSymbol(keywords));
        } else if (isRegisterBtn) {
            setRegisteredChartAry();
            // setIsRegisterBtn(false)
        } else {
            reduceTicker();
        }
    }, [keywords]);

    return (
        <section>
            {!isRegisterBtn ? (
                <div>
                    <div style={{ textAlign: 'right' }}>
                        <Button variant="contained" color="primary" onClick={() => setRegisteredTickers()}>
                            登録済みの銘柄を検索する
                        </Button>
                    </div>
                    <SpaceRow height={10} />
                    <TagSearch setKeyword={setKeywords} isRegisterBtn={isRegisterBtn} />
                </div>
            ) : (
                <div style={{ textAlign: 'right' }}>
                    <Button
                        variant="outlined"
                        color="primary"
                        onClick={() => {
                            setIsRegisterBtn(false);
                            setChartAry([]);
                        }}
                    >
                        銘柄を検索する
                    </Button>
                </div>
            )}
            {chartAry.length ? <StockChart chartAry={chartAry} title={'Compare Chart '} /> : '銘柄を検索できます。'}
        </section>
    );
};

export default Comparison;
