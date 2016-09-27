/* @flow */

import {makeRequest} from './apiHandler.js';

export async function getRandom():Promise<number> {
	const response = await makeRequest('/v1/test/random','GET');
	return response.result;
}
