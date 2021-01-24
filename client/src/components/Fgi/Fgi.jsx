import React, { useEffect, useState } from 'react';
import axios from 'axios';
import FgiChart from './FgiChart';
import FgiTable from './FgiTable';

const Fgi = () => {
    const [fgis, setFgis] = useState([]);
    const dates = fgis
        .map((f) => {
            const date = new Date(f.created_at);
            const day = date.getDate();
            const month = date.getMonth() + 1;
            const fmtDate = month + '/' + day;
            return fmtDate;
        })
        .reverse();

    const nowValues = fgis.map((f) => f.now_value);

    useEffect(() => {
        async function fetchFgis() {
            try {
                const response = await axios.get('/fgi?limit=30');
                setFgis(response.data);
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
