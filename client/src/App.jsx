import React, { useEffect, useState } from 'react'

const App = () => {
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

    console.log(data)
    return (
        <div>
            client
        </div>
    );
}

export default App;
