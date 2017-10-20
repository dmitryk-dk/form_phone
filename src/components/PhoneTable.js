import React, {Component} from 'react';

class PhoneTable extends Component {

    render() {
        return (
            <table>
                <tbody>
                { this.tableBody() }
                </tbody>
            </table>
        );
    }

    tableBody () {
        const {phoneNumbers} = this.props;
        return (
            phoneNumbers.map(phoneNumber => (
                <tr>
                    <td>
                        {phoneNumber.number}
                    </td>
                    <td>
                        <button className="app-button app-button_del">Delete Number</button>
                    </td>
                </tr>
            ))
        );
    };
}

export default PhoneTable;
