export const dateParse = (date) => {
    const newDate = new Date(date);
    const day = newDate.getDate();
    const month = newDate.getMonth() + 1;
    const fmtDate = month + '/' + day;
    return fmtDate;
};
