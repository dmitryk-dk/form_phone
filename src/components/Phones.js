import React from 'react';
import PhoneTable from './PhoneTable';
import Form from './Form';

export default ({
    phoneNumbers,
    submit
}) => {
    return (
        <div className='app-wrapper'>
            <h1>Please enter phone number</h1>
            <Form submit={submit}/>
            <PhoneTable phoneNumbers={phoneNumbers}/>
        </div>
    );
}
