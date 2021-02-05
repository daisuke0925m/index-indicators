import React from 'react';
import Auth from '../../Auth';
import Comparison from '../Comparison/Comparison';
import { CntWrap, SpaceRow } from '../UiKits';
import Fgi from '../Fgi/Fgi';
import FgiDes from '../Fgi/FgiDes';
import { Grid } from '@material-ui/core';
import LikeSwitch from '../Likes/LikeSwitch';
import Skew from '../Skew/Skew';

const Main = () => {
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
                        <LikeSwitch />
                        <Fgi />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap title={'SKEW'} description={<br />} accordionHead={''}>
                    <div>
                        <LikeSwitch />
                        <Skew />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <Auth>
                    <CntWrap title={'Comparison'} description={<br />} accordionHead={''}>
                        <Comparison />
                    </CntWrap>
                </Auth>
            </Grid>
        </Grid>
    );
};

export default Main;
