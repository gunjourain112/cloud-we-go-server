import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 500 },  // 30초 내에 500명 도달
    { duration: '3m', target: 1000 },  // 3분간 1000명까지 늘리고 유지
    { duration: '30s', target: 0 },
  ],
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

export default function () {
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const loginRes = http.post(`${BASE_URL}/auth/login`, JSON.stringify({
    email: 'user1@example.com',
    password: 'password123',
  }), params);
  
  check(loginRes, { 'login success': (r) => r.status === 200 });
  const token = loginRes.json('access_token');

  if (token) {
    const authParams = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    };
    
    const createRes = http.post(`${BASE_URL}/posts`, JSON.stringify({
      title: 'Benchmark Post',
      body: 'Testing performance with JWT and DB write',
      tags: ['bench', 'test'],
    }), authParams);
    
    check(createRes, { 'create post success': (r) => r.status === 201 });
  }

  const postsRes = http.get(`${BASE_URL}/posts?limit=10`);
  check(postsRes, { 'list posts success': (r) => r.status === 200 });

  sleep(1);
}
