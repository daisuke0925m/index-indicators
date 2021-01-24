import React from 'react';
import PropTypes from 'prop-types';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';

const FgiChart = (props) => {
    FgiChart.propTypes = {
        dates: PropTypes.arrayOf(PropTypes.string),
        nowValues: PropTypes.arrayOf(PropTypes.number),
    };

    const options = {
        title: {
            text: 'FGI',
        },
        series: [
            {
                data: props.nowValues,
            },
        ],
        legend: {
            align: 'top',
            verticalAlign: 'left',
            x: 10,
            y: -28,
            itemStyle: {
                color: 'red',
            },
        },
        yAxis: {
            title: {
                text: null,
            },
            gridLineWidth: 1,
            gridLineDashStyle: 'ShortDash',
            gridLineColor: '#A0A0A0',
        },
        xAxis: {
            categories: props.dates,
            type: 'datetime',
            gridLineColor: '#A0A0A0',
            gridLineDashStyle: 'ShortDash',
        },
    };

    return <HighchartsReact highcharts={Highcharts} options={options} />;
};

export default FgiChart;
