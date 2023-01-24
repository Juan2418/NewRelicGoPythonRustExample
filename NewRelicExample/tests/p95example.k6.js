import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
	scenarios: {
		constant_request_rate: {
			executor: 'constant-arrival-rate',
			rate: 1200,
			timeUnit: '1m',
			duration: '2m',
			preAllocatedVUs: 200,
			maxVUs: 200,
		},
	},
	thresholds: {
		http_req_duration: ['p(95)<300'], 
	},
};

const BASE_URL = 'https://new-relic-go-example.fly.dev';
const ENDPOINT = '';

export default () => {
	const authHeaders = {
		"message": "Hello World"
	};

	const objects = http.post(`${BASE_URL}`, authHeaders).json();
	check(objects, {
		'retrieved': (obj) => obj.message,
	});

	sleep(1);
};