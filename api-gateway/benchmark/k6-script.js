import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '20s', target: 6000 },
    { duration: '35s', target: 12000 },
    { duration: '10s', target: 0 },
  ],
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

export default function () {
  const res = http.get(`${BASE_URL}/ping`);
  check(res, { 'ping success': (r) => r.status === 200 });
  sleep(0.1);
}
