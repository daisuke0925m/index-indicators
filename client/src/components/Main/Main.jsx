import React from 'react';
import Fgi from '../Fgi/Fgi';
import FgiDes from '../Fgi/FgiDes';
import { Grid } from '@material-ui/core';
import { CntWrap, SpaceRow, SwitchButton } from '../UiKits';
import Skew from '../Skew/Skew';
import Comparison from '../Comparison/Comparison';
import Auth from '../../Auth';

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
                        <SwitchButton />
                        <Fgi />
                    </div>
                </CntWrap>
            </Grid>
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap title={'SKEW'} description={<br />} accordionHead={''}>
                    <Skew />
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
