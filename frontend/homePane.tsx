import * as React from 'react';
import Paper from 'material-ui/Paper';
import { Card, CardActions, CardHeader, CardTitle, CardText } from 'material-ui/Card';
import CircularProgress from 'material-ui/CircularProgress';

import { getRandom } from './testApi';
require('./homePane.css');

type State = {
	randomNumber: number,
};

export default class HomePane extends React.Component {
	state: State = { randomNumber: 0 };
	render() {
		const { randomNumber } = this.state;
		return (
			<div style={{ flex: 1, display: 'flex' }}>
				<Paper className='paper-pane' zDepth={5}>
					<Card style={{ flex: 1 }}>
						<CardHeader
							title='template webapp'
							avatar={<i className='fa fa-comment-o' aria-hidden='true'></i>}
						/>
						<CardText>
							<p>
								This is where content would go, IF I HAD ANY
							</p>
							<span style={{ margin: '20px' }}>
								a random number:
									{randomNumber}
							</span>
						</CardText>
					</Card>
				</Paper>
			</div>
		);
	}

	componentDidMount() {
		this._getRandomNumber();
	}
	async _getRandomNumber() {
		const x = await getRandom();
		this.setState({ randomNumber: x });
	}
}
