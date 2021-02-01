import httpClient from '../../axios';
import { alertOpenAction } from '../uiState/actions';
import { signInAction, signOutAction } from './actions';

export const signIn = (email, password) => {
    if (email === '' || password === '') {
        return async (dispatch) => {
            try {
                await dispatch(signOutAction());
                await dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'error',
                            message: 'Please Fill all forms!',
                        },
                    })
                );
            } catch (error) {
                console.error(error);
            }
        };
    }
    if (email && password) {
        return async (dispatch) => {
            try {
                await httpClient.post('/login', {
                    email: email,
                    password: password,
                });
                dispatch(
                    signInAction({
                        isSignedIn: true,
                    })
                );
                dispatch(
                    alertOpenAction({
                        alert: {
                            isOpen: true,
                            type: 'success',
                            message: 'signed in',
                        },
                    })
                );
            } catch (error) {
                if (error.response.status == 404) {
                    dispatch(
                        alertOpenAction({
                            alert: {
                                isOpen: true,
                                type: 'error',
                                message: 'User not found',
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

export const listenAuthState = () => {
    return async (dispatch) => {
        try {
            await httpClient.post('/refresh_token');
            dispatch(
                signInAction({
                    isSignedIn: true,
                })
            );
        } catch (error) {
            console.error(error);
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
                        message: 'signed out',
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
                        message: 'Please fill all forms!',
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
                        message: 'Password is not matched',
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
                            message: 'Created User! Please Sign In .',
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
                            message: 'User already exists . Conflict User Name or Email .',
                        },
                    })
                );
            }
            return;
        }
    };
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
