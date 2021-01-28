import React from 'react';
import Highcharts from 'highcharts/highstock';
import HighchartsReact from 'highcharts-react-official';
import PropTypes from 'prop-types';

const StockChart = (props) => {
    StockChart.propTypes = {
        chartAry: PropTypes.arrayOf(
            PropTypes.arrayOf(
                PropTypes.shape({
                    id: PropTypes.number,
                    symbol: PropTypes.string,
                    date: PropTypes.string,
                    open: PropTypes.number,
                    high: PropTypes.number,
                    low: PropTypes.number,
                    close: PropTypes.number,
                    volume: PropTypes.number,
                    created_at: PropTypes.string,
                })
            )
        ),
        title: PropTypes.string,
        chartColor: PropTypes.string,
    };

    const chartColor = props.chartColor;
    const title = props.title;
    const series = props.chartAry.map((ary) => {
        const data = ary.map((d) => {
            return [Date.parse(d.date), d.close];
        });
        const symbol = ary[0].symbol;
        return { name: symbol, data: data };
    });

    const options = {
        chart: {
            spacingRight: 0,
        },
        rangeSelector: {
            buttons: [
                {
                    type: 'day',
                    count: 3,
                    text: '3d',
                },
                {
                    type: 'week',
                    count: 1,
                    text: '1w',
                },
                {
                    type: 'month',
                    count: 1,
                    text: '1m',
                },
                {
                    type: 'month',
                    count: 6,
                    text: '6m',
                },
                {
                    type: 'year',
                    count: 1,
                    text: '1y',
                },
                {
                    type: 'all',
                    text: 'All',
                },
            ],
            selected: 6,
        },
        xAxis: {
            type: 'datetime',
        },
        yAxis: {
            offset: 25,
        },
        title: {
            text: title ? title : null,
        },
        plotOptions: {
            series: {
                lineWidth: 1,
                color: chartColor && chartColor,
            },
        },
        series: series,
    };

    return <HighchartsReact highcharts={Highcharts} constructorType={'stockChart'} options={options} />;
};

export default StockChart;
