/* @flow */

import React from 'react';
import Paper from 'material-ui/Paper';
import {List, ListItem} from 'material-ui/List';
import Divider from 'material-ui/Divider';

const style = {
  width: '200px',
	zDepth: 5
};

type Props = {
	switchPane: (pane: PaneType) => void
}
type State = {}
export default class NavBar extends React.Component {
	state : State
	render() {
		const {switchPane} = this.props;
		return (
			<Paper style={style}>
				<List>
					<ListItem primaryText="Home" onClick={() => switchPane('home')}/>
					<ListItem primaryText="Void" onClick={() => switchPane('void')}/>
				</List>
			</Paper>
		);
	}
}
