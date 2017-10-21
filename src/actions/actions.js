import appDispatcher from '../utils/dispatcher';
import * as http from '../utils/http';
import * as actionTypes from "./types";
import * as consts from '../utils/consts';

export const submit = (phoneNumber) => {
    const url =  consts.host.replace('{API}', consts.postPhone);
    appDispatcher.dispatch({
        type: actionTypes.PHONE_SAVE_REQUSET,
    });
    http.post(url, phoneNumber)
        .then(data => appDispatcher.dispatch({
            type: actionTypes.PHONE_SAVE_SUCCESS,
            phoneNumber
        }))
        .catch(error => appDispatcher.dispatch({
            type: actionTypes.PHONE_SAVE_FAILED,
            error
        }))
};

export const getData = (url) => {
    http.get(url)
        .then(data => appDispatcher.dispatch({
            type: actionTypes.APP_INIT,
            data
        }))
        .catch(error => appDispatcher.dispatch({
            type: actionTypes.APP_INIT_ERROR,
            error
        }))
};
