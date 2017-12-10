import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as injectTapEventPlugin from 'react-tap-event-plugin';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import AppBar from 'material-ui/AppBar';
import Paper from 'material-ui/Paper';
import HomePane from './homePane';

require('./index.css');

window.onload = () => {
	injectTapEventPlugin();
	const content = document.getElementById('content');
	ReactDOM.render(<Index />, content);
	console.log('initialized');
};

type State = {
	pane: PaneType,
	sideBarOpen: boolean
};
export default class Index extends React.Component {
	state: State = {
		pane: 'home',
		sideBarOpen: false
	};
	render() {
		const { pane, sideBarOpen } = this.state;
		let paneElement = <div />;
		switch (pane) {
			case 'home':
				paneElement = <HomePane />;
				break;
			default:
				paneElement = (<h1>Not implemented!</h1>);
		}

		// TODO put navigation in AppBar
		return (
			// muiTheme={getMuiTheme(darkBaseTheme)}
			<MuiThemeProvider>
				<div id='indexRoot'>
					<AppBar title='webapp template' showMenuIconButton={false} />
					<div id='content-view'>
						{paneElement}
					</div>
				</div>
			</MuiThemeProvider >
		);
	}

	_paneSwitch(pane: PaneType) {
		console.log('switching panes', pane);
		this.setState({ pane });
	}
}
