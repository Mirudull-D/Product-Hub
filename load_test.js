import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
  stages: [
    { duration: '30s', target: 70 },
    { duration: '1m', target: 100 },
    { duration: '30s', target: 50 },
  ],
  thresholds: {
    http_req_duration: ['p(95)<1000'],
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080/api/v1';

export default function () {
  const email = `user_${randomString(8)}@example.com`;
  const password = 'password123';

  // Register
  const registerPayload = JSON.stringify({
    firstName: 'Test',
    lastName: 'User',
    email,
    password,
  });

  const regRes = http.post(
    `${BASE_URL}/register`,
    registerPayload,
    {
      headers: {
        'Content-Type': 'application/json',
      },
    }
  );

  check(regRes, {
    'register status 201': (r) => r.status === 201,
  });

  // Login
  const loginPayload = JSON.stringify({
    email,
    password,
  });

  const loginRes = http.post(
    `${BASE_URL}/login`,
    loginPayload,
    {
      headers: {
        'Content-Type': 'application/json',
      },
    }
  );

  const loginSuccess = check(loginRes, {
    'login status ok': (r) => r.status === 200 || r.status === 201,
    'token exists': (r) => {
      try {
        return !!r.json().token;
      } catch {
        return false;
      }
    },
  });

  if (!loginSuccess) {
    console.log('Login failed');
    console.log(loginRes.body);
    return;
  }

  const token = loginRes.json().token;

  // Products
  const productRes = http.get(
    `${BASE_URL}/product`
  );

  check(productRes, {
    'products status 200': (r) => r.status === 200,
  });

  // Checkout
  const checkoutPayload = JSON.stringify({
    items: [
      {
        productId: 1,
        quantity: 1,
      },
    ],
  });

  const checkoutRes = http.post(
    `${BASE_URL}/cart/checkout`,
    checkoutPayload,
    {
      headers: {
        'Content-Type': 'application/json',
        Authorization: token,
      },
    }
  );

  check(checkoutRes, {
    'checkout success': (r) =>
      r.status === 201 || r.status === 400,
  });

  sleep(1);
}