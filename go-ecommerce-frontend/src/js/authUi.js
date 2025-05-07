// src/js/authUi.js
import { loginUser, registerUser } from '../api/auth.js';
import { setToken, setUserRole, displayMessage } from './utils.js';

export function initLoginPage() {
    const loginForm = document.getElementById('login-form');
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const email = e.target.email.value;
            const password = e.target.password.value;
            const result = await loginUser(email, password);

            if (result.token) {
                setToken(result.token);
                setUserRole(result.role); // Assuming backend returns role
                displayMessage('Login successful!', 'success');
                window.updateNav(); // Update nav defined in main.js
                setTimeout(() => window.location.href = '/', 1000);
            } else {
                displayMessage(result.error || 'Login failed.', 'danger');
            }
        });
    }
}

export function initRegisterPage() {
    const registerForm = document.getElementById('register-form');
    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const firstName = e.target.firstName.value;
            const lastName = e.target.lastName.value;
            const email = e.target.email.value;
            const password = e.target.password.value;
            const role = e.target.role.value; // Get role from select

            const result = await registerUser(firstName, lastName, email, password, role);

            if (result.token) {
                setToken(result.token);
                setUserRole(result.role);
                displayMessage('Registration successful! Logging you in...', 'success');
                window.updateNav();
                setTimeout(() => window.location.href = '/', 1000);
            } else {
                displayMessage(result.error || 'Registration failed.', 'danger');
            }
        });
    }
}