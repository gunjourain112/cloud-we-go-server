import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 5000 },
    { duration: '1m', target: 10000 },
    { duration: '30s', target: 0 },
  ],
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

export default function () {
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const createRes = http.post(`${BASE_URL}/posts`, JSON.stringify({
    title: 'Benchmark Post',
    body: 'Raw performance testing without auth',
    tags: ['bench', 'noauth'],
  }), params);
  check(createRes, { 'create post success': (r) => r.status === 201 });

  const postsRes = http.get(`${BASE_URL}/posts?limit=10`);
  check(postsRes, { 'list posts success': (r) => r.status === 200 });

  sleep(1);
}
