import React from 'react';



function Button({text, color, handler}) {
    return (
        <div>
            <button className={color} onClick={handler}>{text}</button>
        </div>
    );
}

export default Button;