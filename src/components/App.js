import React, { Component } from 'react';
import { Container } from 'flux/utils';
import PhoneTable from './PhoneTable';
import Form from './Form';
import appDispatcher from '../utils/dispatcher';

// store
import AppStore from '../stores/store';
const appStore = new AppStore(appDispatcher);
// actions
import * as actions from '../actions/actions'

class App extends Component {

    static getStores() {
        return [ appStore ];
    }

    static calculateState(prevState, props) {
        return {
            ...appStore.getState()
        }
    }

    render() {
        const {phoneNumbers} = this.state;
        return (
            <div className='app-wrapper'>
                <h1>Please enter phone number</h1>
                <Form actions={actions}/>
                <PhoneTable phoneNumbers={phoneNumbers}/>
            </div>
        );
    }
}

export default new Container.create(App);
