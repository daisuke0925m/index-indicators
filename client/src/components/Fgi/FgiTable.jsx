import React from 'react';
import { checkColor, dateParse } from '../Functions/functions';
import PropTypes from 'prop-types';
import { makeStyles, useTheme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableFooter from '@material-ui/core/TableFooter';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import IconButton from '@material-ui/core/IconButton';
import FirstPageIcon from '@material-ui/icons/FirstPage';
import KeyboardArrowLeft from '@material-ui/icons/KeyboardArrowLeft';
import KeyboardArrowRight from '@material-ui/icons/KeyboardArrowRight';
import LastPageIcon from '@material-ui/icons/LastPage';

const useStyles1 = makeStyles((theme) => ({
    root: {
        flexShrink: 0,
        marginLeft: theme.spacing(2.5),
    },
}));

function TablePaginationActions(props) {
    const classes = useStyles1();
    const theme = useTheme();
    const { count, page, rowsPerPage, onChangePage } = props;

    const handleFirstPageButtonClick = (event) => {
        onChangePage(event, 0);
    };

    const handleBackButtonClick = (event) => {
        onChangePage(event, page - 1);
    };

    const handleNextButtonClick = (event) => {
        onChangePage(event, page + 1);
    };

    const handleLastPageButtonClick = (event) => {
        onChangePage(event, Math.max(0, Math.ceil(count / rowsPerPage) - 1));
    };

    return (
        <div className={classes.root}>
            <IconButton onClick={handleFirstPageButtonClick} disabled={page === 0} aria-label="first page">
                {theme.direction === 'rtl' ? <LastPageIcon /> : <FirstPageIcon />}
            </IconButton>
            <IconButton onClick={handleBackButtonClick} disabled={page === 0} aria-label="previous page">
                {theme.direction === 'rtl' ? <KeyboardArrowRight /> : <KeyboardArrowLeft />}
            </IconButton>
            <IconButton
                onClick={handleNextButtonClick}
                disabled={page >= Math.ceil(count / rowsPerPage) - 1}
                aria-label="next page"
            >
                {theme.direction === 'rtl' ? <KeyboardArrowLeft /> : <KeyboardArrowRight />}
            </IconButton>
            <IconButton
                onClick={handleLastPageButtonClick}
                disabled={page >= Math.ceil(count / rowsPerPage) - 1}
                aria-label="last page"
            >
                {theme.direction === 'rtl' ? <FirstPageIcon /> : <LastPageIcon />}
            </IconButton>
        </div>
    );
}

TablePaginationActions.propTypes = {
    count: PropTypes.number.isRequired,
    onChangePage: PropTypes.func.isRequired,
    page: PropTypes.number.isRequired,
    rowsPerPage: PropTypes.number.isRequired,
};

const useStyles2 = makeStyles({
    tablePaginationRoot: {
        '& .MuiTablePagination-caption': {
            display: 'none',
        },
        '& .MuiSelect-root': {
            display: 'none',
        },
        '& .MuiTablePagination-selectIcon': {
            display: 'none',
        },
    },
});

export default function FgiTable(props) {
    const fgis = props.fgis;
    FgiTable.propTypes = {
        fgis: PropTypes.arrayOf(
            PropTypes.shape({
                id: PropTypes.number,
                created_at: PropTypes.string,
                now_value: PropTypes.number,
                now_text: PropTypes.string,
                pc_value: PropTypes.number,
                pc_text: PropTypes.string,
                one_w_value: PropTypes.number,
                one_w_text: PropTypes.string,
                one_m_value: PropTypes.number,
                one_m_text: PropTypes.string,
                one_y_value: PropTypes.number,
                one_y_text: PropTypes.string,
            })
        ),
    };

    const classes = useStyles2();
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(1);

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="custom pagination table">
                {(rowsPerPage > 0 ? fgis.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage) : fgis).map(
                    (row) => (
                        <TableBody key={row.id}>
                            <TableRow>
                                <TableCell colSpan={3}>{dateParse(row.created_at)}</TableCell>
                            </TableRow>
                            <TableRow>
                                <TableCell component="th" scope="row">
                                    Last day
                                </TableCell>
                                <TableCell align="right" style={{ color: checkColor(row.now_value) }}>
                                    {row.now_value}
                                </TableCell>
                            </TableRow>
                            <TableRow>
                                <TableCell component="th" scope="row">
                                    1 Week Ago
                                </TableCell>
                                <TableCell align="right" style={{ color: checkColor(row.one_w_value) }}>
                                    {row.one_w_value}
                                </TableCell>
                            </TableRow>
                            <TableRow>
                                <TableCell component="th" scope="row">
                                    1 Month Ago
                                </TableCell>
                                <TableCell align="right" style={{ color: checkColor(row.one_m_value) }}>
                                    {row.one_m_value}
                                </TableCell>
                            </TableRow>
                            <TableRow>
                                <TableCell component="th" scope="row">
                                    1 Year Ago
                                </TableCell>
                                <TableCell align="right" style={{ color: checkColor(row.one_y_value) }}>
                                    {row.one_y_value}
                                </TableCell>
                            </TableRow>
                        </TableBody>
                    )
                )}
                {/* {emptyRows > 0 && (
                        <TableRow>
                            <TableCell colSpan={6} />
                        </TableRow>
                    )} */}
                <TableFooter>
                    <TableRow>
                        <TablePagination
                            className={classes.tablePaginationRoot}
                            rowsPerPageOptions={[5, 10, 25, { label: 'All', value: -1 }]}
                            colSpan={3}
                            count={fgis.length}
                            rowsPerPage={rowsPerPage}
                            page={page}
                            SelectProps={{
                                inputProps: { 'aria-label': 'rows per page' },
                                native: true,
                            }}
                            onChangePage={handleChangePage}
                            onChangeRowsPerPage={handleChangeRowsPerPage}
                            ActionsComponent={TablePaginationActions}
                        />
                    </TableRow>
                </TableFooter>
            </Table>
        </TableContainer>
    );
}
