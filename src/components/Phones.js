import React from 'react';
import PhoneTable from './PhoneTable';
import Form from './Form';

export default ({
    phones,
    submit,
    deletePhone
}) => {
    return (
        <div className='app-wrapper'>
            <h1>Please enter phone number</h1>
            <Form submit={submit}/>
            <PhoneTable phones={phones} deletePhone={deletePhone}/>
        </div>
    );
}
