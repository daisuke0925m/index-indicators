import React from "react";
import Fgi from "../Fgi/Fgi";
import FgiDes from "../Fgi/FgiDes";
import { Grid } from "@material-ui/core";
import { CntWrap, SpaceRow } from "../UiKits";

const Main = () => {

    return (
        <Grid
            container
            justify="center"
        >
            <Grid item xs={12} sm={8}>
                <SpaceRow height={30} />
                <CntWrap
                    title={'Fear&Greed Index'}
                    description={<FgiDes />}
                    accordionHead={'What is the Fear & Greed Index?'}
                >
                    <Fgi />
                </CntWrap>
            </Grid>
        </Grid >
    )
};

export default Main