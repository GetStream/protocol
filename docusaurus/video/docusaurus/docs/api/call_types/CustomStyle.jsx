import React from 'react';

// Workaround for hiding SDK tutorial links on the call types shared page
const css = `.theme-admonition {
    display: none
}`;

const CustomStyle = () => <style>{css}</style>;

export default CustomStyle;