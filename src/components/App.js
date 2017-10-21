import React, { Component } from 'react';
import { Container } from 'flux/utils';
import Phones from './Phones';


// store
import appStore from '../stores/store';
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
        return (
            <Phones
                {...this.state}
                {...actions}
            />
        );
    }
}

export default new Container.create(App);
