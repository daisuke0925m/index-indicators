import axios from 'axios';
import httpClient from '../../axios';
import { alertOpenAction } from '../uiState/actions';
import { signInAction, signOutAction } from './actions';

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
                        userName: data.userName,
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
                const res = await httpClient.get(`/users/${id}`);
                const data = res.data;
                dispatch(
                    signInAction({
                        isSignedIn: true,
                        userID: data.id,
                        userName: data.userName,
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
                alertOpenAction(
                    {
                        alert: {
                            isOpen: true,
                            type: 'success',
                            message: 'ユーザーを作成しました。 ログインして下さい。',
                        },
                    },
                    console.log('1')
                )
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
        console.log(password);
        return async (dispatch) => {
            try {
                await httpClient.delete(`/users/${id}`, {
                    data: { password: password },
                });
                dispatch(
                    signInAction({
                        isSignedIn: false,
                    })
                );
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
    } else {
        return;
    }
};

// export const resetPassword = (email) => {
//     return async (dispatch) => {
//         if (email === "") {
//             alert("必須項目が未入力です。")
//             return false
//         } else {
//             auth.sendPasswordResetEmail(email)
//                 .then(() => {
//                     alert('入力されたアドレスにパスワードリセット用のメールを送りました。')
//                     dispatch(push('/signin'))
//                 }).catch(() => {
//                     alert('パスワードリセットに失敗しました。')
//                 })
//         }
//     }
// }
