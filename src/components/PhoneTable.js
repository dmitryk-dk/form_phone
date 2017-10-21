import React, {Component} from 'react';

class PhoneTable extends Component {

    render() {
        return (
            <div className="app-table-wrapper">
                <table>
                    <tbody>
                    { this.tableBody() }
                    </tbody>
                </table>
            </div>

        );
    }

    tableBody () {
        const {phones, deletePhone} = this.props;
        return (
            phones.map((phone, i) => (
                <tr>
                    <td>{i}</td>
                    <td>
                        {phone.number}
                    </td>
                    <td>
                        <button
                            className="app-button app-button_del"
                            onClick={() =>deletePhone(phone)}
                        >Delete Number</button>
                    </td>
                </tr>
            ))
        );
    };
}

export default PhoneTable;
