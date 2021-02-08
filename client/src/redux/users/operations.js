import axios from 'axios';
import httpClient from '../../axios';
import { alertOpenAction } from '../uiState/actions';
import { fetchLikesAction, signInAction, signOutAction } from './actions';

export const signIn = (email, password) => {
    if (email === '' || password === '') {
        return async (dispatch) => {
            try {
                await dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'error',
                            message: '全てのフォームに記入してください。',
                        },
                    })
                );
                return;
            } catch (error) {
                console.error(error);
                return;
            }
        };
    }
    if (email && password) {
        return async (dispatch) => {
            try {
                const res = await httpClient.post('/login', {
                    email: email,
                    password: password,
                });
                const data = res.data;
                dispatch(
                    signInAction({
                        isSignedIn: true,
                        userID: data.id,
                        userName: data.user_name,
                    })
                );
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'success',
                            message: 'ログインしました。',
                        },
                    })
                );
                return;
            } catch (error) {
                if (error.response.status == 404) {
                    dispatch(
                        alertOpenAction({
                            alert: {
                                isOpen: true,
                                type: 'error',
                                message: 'ユーザーが見つかりません。',
                            },
                        })
                    );
                }
                return;
            }
        };
    }
};

const httpClientSingle = axios.create({
    baseURL: 'http://localhost:8080',
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
});

export const listenAuthState = () => {
    return async (dispatch) => {
        try {
            const res = await httpClientSingle.post('/refresh_token');
            const id = res.data.id;
            try {
                const res = await httpClientSingle.get(`/users/${id}`);
                const data = res.data;
                dispatch(
                    signInAction({
                        isSignedIn: true,
                        userID: data.id,
                        userName: data.user_name,
                    })
                );
                return;
            } catch (error) {
                console.error(error);
                return;
            }
        } catch (error) {
            if (error.response.status == 404 || error.response.status == 401) {
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'info',
                            message: '全ての機能を試すにはログインしてください。',
                        },
                    })
                );
            }
            return;
        }
    };
};

export const signOut = () => {
    return async (dispatch) => {
        try {
            await httpClient.post('/logout');
            dispatch(signOutAction());
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'success',
                        message: 'ログアウトしました。',
                    },
                })
            );
        } catch (error) {
            console.error(error);
        }
    };
};

export const signUp = (username, email, password, confirmPassword) => {
    return async (dispatch) => {
        if (username === '' || email === '' || password === '' || confirmPassword === '') {
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'error',
                        message: '全てのフォームに記入してください!',
                    },
                })
            );
            return;
        }

        if (password !== confirmPassword) {
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'error',
                        message: 'パスワードが一致しません！',
                    },
                })
            );
            return;
        }

        try {
            await httpClient.post('/users', {
                user_name: username,
                email: email,
                password: password,
            });
            await dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'success',
                        message: 'ユーザーを作成しました。 ログインして下さい。',
                    },
                })
            );
        } catch (error) {
            if (error.response.status == 409) {
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'error',
                            message: 'ユーザーネームまたは、Eメールはすでに使用されております。',
                        },
                    })
                );
            }
            return;
        }
    };
};

export const deleteUser = (password, id) => {
    if (password === '') {
        return async (dispatch) => {
            try {
                await dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'error',
                            message: 'パスワードを記入して下さい。',
                        },
                    })
                );
                return;
            } catch (error) {
                console.error(error);
                return;
            }
        };
    }
    if (password) {
        return async (dispatch) => {
            try {
                await httpClient.delete(`/users/${id}`, {
                    data: { password: password },
                });
                dispatch(signOutAction());
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'warning',
                            message: 'ユーザーを削除しました。',
                        },
                    })
                );
            } catch (error) {
                if (error.response.status == 404 || error.response.status == 400) {
                    dispatch(
                        alertOpenAction({
                            alert: {
                                isOpen: true,
                                type: 'error',
                                message: 'パスワードが一致しません。 もう一度お試し下さい。',
                            },
                        })
                    );
                }
            }
        };
    }
};

export const updateUser = (newName, newEmail, newPass, pass, id) => {
    if (newName == '' && newEmail == '' && newPass == '') {
        return async (dispatch) => {
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'error',
                        message: '送信できませんでした。少なくとも1つの項目は変更して下さい。',
                    },
                })
            );
            return;
        };
    }

    if (pass) {
        return async (dispatch) => {
            try {
                await httpClient.put(`/users/${id}`, {
                    user: {
                        password: pass,
                    },
                    new_user: {
                        user_name: newName,
                        email: newEmail,
                        password: newPass,
                    },
                });
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'success',
                            message:
                                `${newName.length ? 'ユーザー名 ' : ''}` +
                                `${newEmail.length ? 'Email ' : ''}` +
                                `${newPass.length ? 'パスワード ' : ''}` +
                                'の変更に成功しました。',
                        },
                    })
                );
                dispatch(listenAuthState());
            } catch (error) {
                if (error.response.status == 404 || error.response.status == 400) {
                    dispatch(
                        alertOpenAction({
                            alert: {
                                isOpen: true,
                                type: 'error',
                                message: 'パスワードが一致しません。 もう一度お試し下さい。',
                            },
                        })
                    );
                }
            }
        };
    } else {
        return async (dispatch) => {
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'error',
                        message: 'パスワードを記入してください。',
                    },
                })
            );
            return;
        };
    }
};

export const fetchUsersLikes = (userID) => {
    return async (dispatch) => {
        try {
            const res = await httpClient.get(`/users/${userID}/likes`);
            const likes = res.data.likes;
            if (likes) {
                dispatch(fetchLikesAction(likes));
            }
            return;
        } catch (error) {
            console.error(error);
        }
    };
};

export const likePost = (userID, symbol) => {
    return async (dispatch) => {
        try {
            await httpClient.post(`users/${userID}/likes`, {
                symbol: symbol,
            });
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'success',
                        message: `${symbol}を登録しました。`,
                    },
                })
            );
            dispatch(fetchUsersLikes(userID));
            return;
        } catch (error) {
            console.error(error);
        }
    };
};

export const likeDelete = (userID, likeID, symbol) => {
    return async (dispatch) => {
        try {
            await httpClient.delete(`users/${userID}/likes/${likeID}`);
            dispatch(
                alertOpenAction({
                    alert: {
                        isOpen: true,
                        type: 'warning',
                        message: `${symbol}を解除しました。`,
                    },
                })
            );
            dispatch(fetchUsersLikes(userID));
            return;
        } catch (error) {
            console.error(error);
        }
    };
};
