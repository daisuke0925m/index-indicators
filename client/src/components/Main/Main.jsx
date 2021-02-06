import React, { useState } from 'react';
import Comparison from '../Comparison/Comparison';
import { CntWrap, SpaceRow } from '../UiKits';
import Fgi from '../Fgi/Fgi';
import FgiDes from '../Fgi/FgiDes';
import { Grid } from '@material-ui/core';
import { getUsersLikes } from '../../redux/users/selectors';
import LikeSwitch from '../Likes/LikeSwitch';
import Skew from '../Skew/Skew';
import { useSelector } from 'react-redux';
import { useEffect } from 'react';

const Main = () => {
    const selector = useSelector((state) => state);
    const likes = getUsersLikes(selector);
    const [fgiFlag, setFgiFlag] = useState(false);
    const [skewFlag, setSkewFlag] = useState(false);

    const checkLikes = () => {
        for (let i = 0; i < likes.length; i++) {
            if (likes[i].symbol == 'fgi') {
                setFgiFlag(true);
            }
            if (likes[i].symbol == '^skew') {
                setSkewFlag(true);
            }
        }
    };

    useEffect(() => {
        checkLikes();
    }, [likes]);

    return (
        <Grid container justify="center">
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap
                    title={'Fear&Greed Index'}
                    description={<FgiDes />}
                    accordionHead={'What is the Fear & Greed Index?'}
                >
                    <div>
                        <LikeSwitch flag={fgiFlag} symbol={'fgi'} />
                        <Fgi />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap title={'SKEW'} description={<br />} accordionHead={''}>
                    <div>
                        <LikeSwitch flag={skewFlag} symbol={'^skew'} />
                        <Skew />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap title={'Comparison'} description={<br />} accordionHead={''}>
                    <Comparison />
                </CntWrap>
            </Grid>
        </Grid>
    );
};

export default Main;
