import { Button, IconButton } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import httpClient from '../../axios';
import { getUsersLikes } from '../../redux/users/selectors';
import StockChart from '../Chart/StockChart';
import { SpaceRow, TagSearch } from '../UiKits';
import Skeleton from '@material-ui/lab/Skeleton';
import AutorenewRoundedIcon from '@material-ui/icons/AutorenewRounded';
import SwitchTable from './SwitchTable';
import SwitchPostTable from './SwitchPostTable';
import Auth from '../../Auth';

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
                        <Auth
                            enableEle={
                                <Button variant="contained" color="primary" onClick={() => setRegisteredTickers()}>
                                    登録済みの銘柄を表示する
                                </Button>
                            }
                            disableEle={
                                <Button variant="contained" disabled={true}>
                                    登録済みの銘柄を表示する
                                </Button>
                            }
                        />
                    </div>
                    <SpaceRow height={10} />
                    <TagSearch setKeyword={setKeywords} isRegisterBtn={isRegisterBtn} />
                    <SwitchPostTable keywords={keywords} />
                    {chartAry.length ? (
                        <StockChart chartAry={chartAry} title={'Compare Chart '} />
                    ) : (
                        <div style={{ padding: 20 }}>
                            <span>銘柄を検索できます。</span>
                            <Skeleton variant="rect" width={'100%'} height={200} animation={false} />
                        </div>
                    )}
                </div>
            ) : (
                <div>
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
                        <IconButton
                            style={{ marginLeft: 10, padding: 5 }}
                            color="primary"
                            onClick={() => setRegisteredTickers()}
                        >
                            <AutorenewRoundedIcon />
                        </IconButton>
                    </div>
                    <SwitchTable />
                    {chartAry.length ? (
                        <StockChart chartAry={chartAry} title={'Compare Chart '} />
                    ) : (
                        <div style={{ padding: 20 }}>
                            <span>登録した銘柄を一括表示できます。</span>
                            <Skeleton variant="rect" width={'100%'} height={200} animation={false} />
                        </div>
                    )}
                </div>
            )}
        </section>
    );
};

export default Comparison;
