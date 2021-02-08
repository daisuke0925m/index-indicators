import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { useDispatch, useSelector } from 'react-redux';
import { getUserID, getUsersLikes } from '../../redux/users/selectors';
import PropTypes from 'prop-types';
import { IconButton, Tooltip } from '@material-ui/core';
import PlaylistAddRoundedIcon from '@material-ui/icons/PlaylistAddRounded';
import { likePost } from '../../redux/users/operations';
import { alertOpenAction } from '../../redux/uiState/actions';
import PlaylistAddCheckRoundedIcon from '@material-ui/icons/PlaylistAddCheckRounded';
import Auth from '../../Auth';

const useStyles = makeStyles({
    table: {
        // minWidth: 650,
    },
});

const SwitchPostTable = (props) => {
    SwitchPostTable.propTypes = {
        keywords: PropTypes.arrayOf(PropTypes.string),
    };

    const dispatch = useDispatch();
    const symbols = props.keywords;
    const classes = useStyles();
    const selector = useSelector((state) => state);
    const likes = getUsersLikes(selector);
    const userID = getUserID(selector);

    const checkSymbol = (symbol) => {
        let isLiked = false;
        for (let i = 0; i < likes.length; i++) {
            const likedSymbol = likes[i].symbol;
            if (symbol == likedSymbol) {
                isLiked = true;
            }
        }
        return isLiked;
    };

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
                    {symbols.map((symbol) => (
                        <TableRow key={symbol}>
                            <TableCell component="th" scope="row">
                                {symbol}
                            </TableCell>
                            <TableCell align="right">
                                {!checkSymbol(symbol) ? (
                                    <Auth
                                        enableEle={
                                            <Tooltip title="銘柄を登録">
                                                <IconButton
                                                    color="primary"
                                                    onClick={() => {
                                                        dispatch(likePost(userID, symbol));
                                                    }}
                                                >
                                                    <PlaylistAddRoundedIcon />
                                                </IconButton>
                                            </Tooltip>
                                        }
                                        disableEle={
                                            <div style={{ textAlign: 'right' }}>
                                                <Tooltip title="ログインユーザーのみ登録できます。" aria-label="add">
                                                    <IconButton>
                                                        <PlaylistAddRoundedIcon />
                                                    </IconButton>
                                                </Tooltip>
                                            </div>
                                        }
                                    />
                                ) : (
                                    <Tooltip title="登録済み">
                                        <IconButton
                                            color="primary"
                                            onClick={() => {
                                                dispatch(
                                                    alertOpenAction({
                                                        alert: {
                                                            isOpen: true,
                                                            type: 'warning',
                                                            message: `${symbol}はすでに登録済みです。`,
                                                        },
                                                    })
                                                );
                                            }}
                                        >
                                            <PlaylistAddCheckRoundedIcon />
                                        </IconButton>
                                    </Tooltip>
                                )}
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
};

export default SwitchPostTable;
