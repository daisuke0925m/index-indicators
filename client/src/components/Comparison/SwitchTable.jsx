import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { useSelector } from 'react-redux';
import { getUsersLikes } from '../../redux/users/selectors';
import LikeSwitch from '../Likes/LikeSwitch';

const useStyles = makeStyles({
    table: {
        // minWidth: 650,
    },
});

const SwitchTable = () => {
    const classes = useStyles();
    const selector = useSelector((state) => state);
    const likes = getUsersLikes(selector);

    const rows = likes.reduce((newAry, like) => {
        if (like.symbol !== 'fgi') {
            newAry.push({ symbol: like.symbol, id: like.id });
        }
        return newAry;
    }, []);

    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>symbol</TableCell>
                        <TableCell align="right">登録ボタン</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.id}>
                            <TableCell component="th" scope="row">
                                {row.symbol}
                            </TableCell>
                            <TableCell component="th" scope="row">
                                <LikeSwitch flag={true} symbol={row.symbol} likeID={row.id} />
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
};

export default SwitchTable;
