import React from 'react';
import Input from './Input'

function Form() {
    return (
        <form className="signup-form mdc-theme--dark" method="POST">
        <Input type="text" name="firstname" placeholder="Firstname"/>
        <Input type="text" name="lastname" placeholder="Lastname"/>
        <Input type="email" name="email" placeholder="Email"/>
        <Input type="password" name="password" placeholder="Password"/>
        <button className="mdc-button mdc-button--primary mdc-button--raised">Sign up </button>
        </form>
    );
}

export default Form;