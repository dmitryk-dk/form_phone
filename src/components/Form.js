import React, {Component} from 'react';

class Form extends Component {

    constructor () {
        super();
        this.state={
            phoneNumber: ''
        };
        this.change = this.change.bind(this);
    }

    render() {
        const {submit} = this.props;
        const {phoneNumber} = this.state;
        return (
            <form onSubmit={() => submit(phoneNumber)}>
                <label htmlFor="phone">
                    <input
                        type="text"
                        value={phoneNumber}
                        id="phone"
                        onChange={this.change}
                    />
                </label>
                <input
                    type="submit"
                    value="Enter"
                    className="app-button app-button_green"
                />
                <small className="help-text">Please use format phone number as 0675675678</small>
            </form>
        );
    }

    change (event) {
        this.setState({phoneNumber: event.target.value});
    }
}

export default Form;
