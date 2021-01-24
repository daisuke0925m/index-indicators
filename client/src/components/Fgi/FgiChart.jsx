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
        legend: {
            enabled: false,
        },
        yAxis: {
            title: {
                text: null,
            },
            gridLineWidth: 1,
            gridLineDashStyle: 'ShortDash',
            gridLineColor: '#A0A0A0',
            max: 100,
        },
        xAxis: {
            title: {
                text: null,
            },
            categories: props.dates,
            type: 'datetime',
            gridLineColor: '#A0A0A0',
            gridLineDashStyle: 'ShortDash',
        },
        plotOptions: {
            pointWidth: 0,
        },
        series: [
            {
                type: 'area',
                data: props.nowValues,
                lineWidth: 0,
                marker: {
                    enabled: false,
                },
                fillColor: {
                    linearGradient: {
                        y1: 1,
                        y2: 0,
                        x1: 0,
                        x2: 0,
                    },
                    stops: [
                        [0, '#F44545'],
                        [0.1, '#F26969'],
                        [0.5, '#ECF07D'],
                        [0.8, '#18BA8F'],
                    ],
                },
            },
        ],
    };

    return <HighchartsReact highcharts={Highcharts} options={options} />;
};

export default FgiChart;
