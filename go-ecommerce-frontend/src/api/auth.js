// src/api/auth.js
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function loginUser(email, password) {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
    });
    return response.json();
}

export async function registerUser(firstName, lastName, email, password, role = 'customer') {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ first_name: firstName, last_name: lastName, email, password, role }),
    });
    return response.json();
}

export async function getUserProfile(token) {
    const response = await fetch(`${API_BASE_URL}/me`, {
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });
    return response.json();
}