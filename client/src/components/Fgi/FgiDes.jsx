import React from 'react';
import BlockQuote from '../UiKits/BlockQuote';

const ComponentName = () => {
    return (
        <BlockQuote>
            <div>
                Investors are driven by two emotions: fear and greed. Too much fear can sink stocks well below where
                they should be. When investors get greedy, they can bid up stock prices way too far.
                <br />
                So what emotion is driving the market now? CNNMoney{"'"}s Fear & Greed index makes it clear.
                <br />
                We look at 7 indicators:
                <br />
                •Stock Price Momentum: The S&P 500 (SPX) versus its 125-day moving average
                <br />
                •Stock Price Strength: The number of stocks hitting 52-week highs and lows on the New York Stock
                Exchange
                <br />
                •Stock Price Breadth: The volume of shares trading in stocks on the rise versus those declining.
                <br />
                •Put and Call Options: The put/call ratio, which compares the trading volume of bullish call options
                relative to the trading volume of bearish put options
                <br />
                •Junk Bond Demand: The spread between yields on investment grade bonds and junk bonds
                <br />
                •Market Volatility: The VIX (VIX), which measures volatility
                <br />
                •Safe Haven Demand: The difference in returns for stocks versus Treasuries
                <br />
                For each indicator, we look at how far they{"'"}ve veered from their average relative to how far they
                normally veer. We look at each on a scale from 0 - 100. The higher the reading, the greedier investors
                are being, and 50 is neutral.
                <br />
                Then we put all the indicators together - equally weighted - for a final index reading.
                <br />
                When the S&P 500 (SPX) plummeted to a three-year low on Sept. 17, 2008 - the height of the financial
                crisis -- the Fear and Greed index sank to 12. The index gained some ground to 28 before stocks finally
                bottomed out on March 9, 2009 and the latest bull market began.
                <br />
                Most recently, in the first quarter of 2012, stocks staged their best run in decades, and the index
                showed pure greed.
                <br />
                <p style={{ textAlign: 'right' }}>from cnn Money</p>
            </div>
        </BlockQuote>
    );
};

export default ComponentName;
