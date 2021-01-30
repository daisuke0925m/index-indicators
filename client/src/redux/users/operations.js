import httpClient from '../../axios';
import { signInAction, signOutAction } from './actions';

export const signIn = (email, password) => {
    // TODOバリデーション
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

// export const signUp = (username, email, password, confirmPassword) => {
//     return async (dispatch) => {
//         if (username === "" || email === "" || password === "" || confirmPassword === "") {
//             alert("必須項目が未入力です")
//             return false
//         }

//         if (password !== confirmPassword) {
//             alert("パスワードが一致しません。もう一度お試しください。")
//             return false
//         }

//         return auth.createUserWithEmailAndPassword(email, password)
//             .then(result => {
//                 const user = result.user

//                 if (user) {
//                     const uid = user.uid
//                     const timestamp = FirebaseTimestamp.now()
//                     const userInitialData = {
//                         created_at: timestamp,
//                         email: email,
//                         role: "customer",
//                         uid: uid,
//                         updated_at: timestamp,
//                         username: username
//                     }

//                     db.collection('users').doc(uid).set(userInitialData)
//                         .then(() => {
//                             dispatch(push('/'))
//                         })
//                 }
//             })
//     }
// }

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
