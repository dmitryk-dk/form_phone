import appDispatcher from '../utils/dispatcher';
import * as http from '../utils/http';
import {APP_INIT, APP_INIT_ERROR} from "./types";
import * as consts from '../utils/consts';

export const submit = (phoneNumber) => {
    const url =  consts.host.replace('{API}', consts.postPhone);
    appDispatcher.dispatch({
        type: '',
    });
    http.post(url, phoneNumber)
        .then(data => console.log(data))
        .catch(reason => console.log(reason.message))
};

export const getData = (url) => {
    http.get(url)
        .then(data => appDispatcher.dispatch({
            type: APP_INIT,
            data
        }))
        .catch(error => appDispatcher.dispatch({
            type: APP_INIT_ERROR,
            error
        }))
};
