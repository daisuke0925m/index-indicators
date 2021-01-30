import React, { useEffect, useState } from 'react';
import { dateParse } from '../Functions/functions';
import FgiChart from './FgiChart';
import FgiTable from './FgiTable';
import httpClient from '../../axios';

const Fgi = () => {
    const [fgis, setFgis] = useState([]);
    const dates = fgis
        .map((f) => {
            const fmtDate = dateParse(f.created_at);
            return fmtDate;
        })
        .reverse();

    const nowValues = fgis.map((f) => f.now_value);

    useEffect(() => {
        async function fetchFgis() {
            try {
                const response = await httpClient.get('/fgi?limit=30');
                setFgis(response.data.fgis);
            } catch (error) {
                console.log(error);
                setFgis([]);
            }
        }
        fetchFgis();
    }, []);

    return (
        <div>
            {fgis.length ? (
                <div>
                    <FgiChart nowValues={nowValues} dates={dates} />
                    <FgiTable fgis={fgis} />
                </div>
            ) : (
                'loading'
            )}
        </div>
    );
};

export default Fgi;
