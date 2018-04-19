import React from 'react';

function Input({type, name, placeholder}) {
    return (
        <div className="mdc-textfield">
        <input
         type={type} 
         name={name}
         placeholder={placeholder}
         className="mdc-textfield__input"/>
       </div>
    );
}

export default Input;