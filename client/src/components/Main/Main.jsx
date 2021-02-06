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
    const [isLikedFgi, setIsLikedFgi] = useState({ isFgi: false, id: 0 });
    const [isLikedSkew, setIsLikedSkew] = useState({ isSkew: false, id: 0 });

    const checkLikes = () => {
        for (let i = 0; i < likes.length; i++) {
            if (likes[i].symbol == 'fgi') {
                setIsLikedFgi({ isFgi: true, id: likes[i].id });
            }
            if (likes[i].symbol == '^skew') {
                setIsLikedSkew({ isSkew: true, id: likes[i].id });
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
                        <LikeSwitch flag={isLikedFgi.isFgi} symbol={'fgi'} likeID={isLikedFgi.id} />
                        <Fgi />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap title={'SKEW'} description={<br />} accordionHead={''}>
                    <div>
                        <LikeSwitch flag={isLikedSkew.isSkew} symbol={'^skew'} likeID={isLikedSkew.id} />
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
