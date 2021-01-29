import { signInAction } from './actions';
import axios from 'axios';

export const signIn = (email, password) => {
    return async (dispatch) => {
        try {
            const response = await axios.post('/login', {
                email: email,
                password: password,
            });
            const data = response.data;
            console.log(data);
            dispatch(
                signInAction({
                    isSignedIn: true,
                })
            );
        } catch (error) {
            console.log(error);
        }
    };
};

// export const signOut = (accessToken) => {
//     return async (dispatch) => {
//         try {
//             const response = await axios.post('/logout', {
//                 headers: {
//                     'authorization': `Bearer ${accessToken}`,
//                 }
//             });
//             const data = response.data
//             dispatch(signOutAction());
//         } catch (error) {
//             console.log(error);
//         }
//     }
// }

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
