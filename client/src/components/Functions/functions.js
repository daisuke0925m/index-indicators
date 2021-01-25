export const dateParse = (date) => {
    const newDate = new Date(date);
    const day = newDate.getDate();
    const month = newDate.getMonth() + 1;
    const fmtDate = month + '/' + day;
    return fmtDate;
};

export const tickerDateParse = (date, num) => {
    date.setDate(date.getDate() + num);
    const fmtDate =
        date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);
    return fmtDate;
};

export const checkColor = (value) => {
    if (value >= 70) {
        const color = '#18BA8F';
        return color;
    }
    if (value <= 30) {
        const color = '#F44545';
        return color;
    }
};
