// src/js/main.js
import { getToken, removeToken, getUserRole, removeUserRole, updateCartCount } from './utils.js';
import { initLoginPage } from './authUI.js';
import { initRegisterPage } from './authUI.js';
import { loadProducts, initProductDetailPage } from './productUi.js';
import { initCartPage } from './cartUi.js';
import { initAdminProductsPage, initAddProductPage, initEditProductPage } from './adminproductUi.js';
import { initProfilePage } from './profileUi.js';
// import { initCheckoutPage } from './checkoutUi.js'; // Placeholder for checkout

function updateNav() {
    const token = getToken();
    const role = getUserRole();
    const guestLinks = document.getElementById('guest-links');
    const userLinks = document.getElementById('user-links');
    const adminLinks = document.getElementById('admin-links');
    const logoutLink = document.getElementById('logout-link');

    if (token) {
        guestLinks.style.display = 'none';
        userLinks.style.display = 'flex'; // Use flex for li items
        logoutLink.style.display = 'block';
        if (role === 'admin') {
            adminLinks.style.display = 'flex'; // Use flex for li items
        } else {
            adminLinks.style.display = 'none';
        }
    } else {
        guestLinks.style.display = 'flex'; // Use flex for li items
        userLinks.style.display = 'none';
        adminLinks.style.display = 'none';
        logoutLink.style.display = 'none';
    }
    updateCartCount();
}

function handleLogout() {
    removeToken();
    removeUserRole();
    updateNav();
    window.location.href = '/'; // Redirect to home
}

document.addEventListener('DOMContentLoaded', () => {
    updateNav();

    const logoutButton = document.getElementById('logout-button');
    if (logoutButton) {
        logoutButton.addEventListener('click', handleLogout);
    }

    // Page-specific initializations
    const path = window.location.pathname;

    if (path === '/' || path === '/index.html') {
        loadProducts();
    } else if (path.startsWith('/src/pages/login.html')) {
        initLoginPage();
    } else if (path.startsWith('/src/pages/register.html')) {
        initRegisterPage();
    } else if (path.startsWith('/src/pages/product-detail.html')) {
        initProductDetailPage();
    } else if (path.startsWith('/src/pages/cart.html')) {
        initCartPage();
    } else if (path.startsWith('/src/pages/profile.html')) {
        initProfilePage();
    }
    // else if (path.startsWith('/src/pages/checkout.html')) {
    //     initCheckoutPage(); // Placeholder
    // }
    else if (path.startsWith('/src/pages/admin/products.html')) {
        initAdminProductsPage();
    } else if (path.startsWith('/src/pages/admin/add-product.html')) {
        initAddProductPage();
    } else if (path.startsWith('/src/pages/admin/edit-product.html')) {
        initEditProductPage();
    }
});

// Expose updateNav globally if needed by other modules, or pass as callback
window.updateNav = updateNav;