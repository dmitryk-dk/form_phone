// utils
import { ReduceStore } from 'flux/utils';
// consts
import { APP_INIT } from '../actions/types';


export default class AppStore extends ReduceStore {

    getInitialState () {
        return {
            errors: null,
            phoneNumbers: [],
        };
    }

    reduce (state, action) {
        switch (action.type) {
            case APP_INIT:
                return {
                    ...state,
                    phoneNumbers: action.data
                };

            default:
                return state;
        };
    };
}


