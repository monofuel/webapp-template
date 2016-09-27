/* @flow */

require("babel-core/register");
require("babel-polyfill");

import Index from './index.jsx';
import React from 'react';
import ReactDOM from 'react-dom';

import injectTapEventPlugin from 'react-tap-event-plugin';

// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();

window.onload = () => {
	const content = document.getElementById('content');
	ReactDOM.render(<Index/>,content);
	console.log('initialized');
}
