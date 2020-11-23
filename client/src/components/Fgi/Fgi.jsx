import React, { useEffect, useState } from 'react'

const Fgi = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        async function fetchData() {
            const resResult = await fetch('http://localhost:8080/api/fgi/')
            resResult
                .json()
                .then((resResult) => setData(resResult))
                .catch(() => null)
        }
        fetchData()
    }, []);

    return (
        <div>
            {
                data.length ? data.map(d => (
                    <div>
                        {d.now_value}
                    </div>
                ))
                    : ('loading')
            }
        </div>
    );
}

export default Fgi;
