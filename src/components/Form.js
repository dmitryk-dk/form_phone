import React, {Component} from 'react';
import {validate} from "../utils/validator";

class Form extends Component {

    constructor (props) {
        super(props);
        this.state={
            number: '',
            hasError: false,
        };
        this.change = this.change.bind(this);
        this.submit = this.submit.bind(this);
    }

    render() {
        const {number, hasError} = this.state;
        return (
            <form onSubmit={(event) => this.submit(event, this.state)}>
                <label htmlFor="phone" className="app-phone_label">
                    <input
                        type="text"
                        value={number}
                        id="phone"
                        className={hasError ? 'app-phone_error': ''}
                        onChange={this.change}
                    />
                </label>
                <input
                    type="submit"
                    value="Add"
                    className="app-button app-button_green"
                />
                {
                    hasError ?
                        <small className="help-text_error">
                            Wrong format number. Example: 0675675678.
                        </small>
                        :
                        <small className="help-text">Please use format phone number as 0675675678</small>
                }

            </form>
        );
    }

    change (event) {
        this.setState({number: event.target.value});
    }

    submit (event, phone) {
        event.preventDefault();
        const {submit} = this.props;
        const {number} = this.state;
        if (validate(number)) {
            this.setState({hasError: false});
            submit(phone);
        } else {
            this.setState({hasError: true});
        }

    }
}

export default Form;
