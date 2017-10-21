import React, {Component} from 'react';

class Form extends Component {

    constructor (props) {
        super(props);
        this.state={
            number: ''
        };
        this.change = this.change.bind(this);
        this.submit = this.submit.bind(this);
    }

    render() {
        const {number} = this.state;
        return (
            <form onSubmit={(event) => this.submit(event, this.state)}>
                <label htmlFor="phone">
                    <input
                        type="text"
                        value={number}
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
        this.setState({number: event.target.value});
    }

    submit (event, phoneNumber) {
        event.preventDefault();
        const {submit} = this.props;
        submit(phoneNumber)
    }
}

export default Form;
