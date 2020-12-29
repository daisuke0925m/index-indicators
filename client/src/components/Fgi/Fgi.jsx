import React, { useEffect, useState } from 'react'

const Fgi = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        async function fetchData() {
            const resResult = await fetch('/server/api/fgi', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    // 'authorization': 'Bearer xxx', 認証してcookieからjwtを取り出す。
                }
            })
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
                data.length ? data.map((d, i) => (
                    <div key={i}>
                        {d.now_value}
                    </div>
                ))
                    : ('loading')
            }
        </div >
    );
}

export default Fgi;
