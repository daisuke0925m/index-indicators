import httpClient from '../../axios';
import { signInAction, signOutAction } from './actions';

export const signIn = (email, password) => {
    // TODOバリデーション
    if (email === '' || password === '') {
        alert('必須項目が未入力です');
        return async (dispatch) => {
            try {
                await dispatch(signOutAction());
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
            } catch (error) {
                console.error(error);
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
            if (error.response.status == 401) {
                console.log('done');
                // TODOエラーハンドリング
            }
        }
    };
};

export const signOut = () => {
    return async (dispatch) => {
        try {
            await httpClient.post('/logout');
            dispatch(signOutAction());
        } catch (error) {
            console.log(error);
        }
    };
};

export const signUp = (username, email, password, confirmPassword) => {
    return async (dispatch) => {
        if (username === '' || email === '' || password === '' || confirmPassword === '') {
            alert('必須項目が未入力です');
            dispatch(signOutAction());
        }

        if (password !== confirmPassword) {
            alert('パスワードが一致しません。もう一度お試しください。');
            dispatch(signOutAction());
        }

        try {
            await httpClient.post('/users', {
                user_name: username,
                email: email,
                password: password,
            });
            dispatch(signOutAction());
        } catch {
            dispatch(signOutAction());
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
