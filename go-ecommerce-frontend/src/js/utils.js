// src/js/utils.js
export function getToken() {
    return localStorage.getItem('authToken');
}

export function setToken(token) {
    localStorage.setItem('authToken', token);
}

export function removeToken() {
    localStorage.removeItem('authToken');
}

export function getUserRole() {
    return localStorage.getItem('userRole');
}

export function setUserRole(role) {
    localStorage.setItem('userRole', role);
}

export function removeUserRole() {
    localStorage.removeItem('userRole');
}


export function getCart() {
    const cart = localStorage.getItem('cart');
    return cart ? JSON.parse(cart) : [];
}

export function saveCart(cart) {
    localStorage.setItem('cart', JSON.stringify(cart));
    updateCartCount(); // Update navbar count whenever cart is saved
}

export function addToCart(product, quantity = 1) {
    const cart = getCart();
    const existingItem = cart.find(item => item.id === product.ID); // Backend product.ID
    if (existingItem) {
        existingItem.quantity += quantity;
    } else {
        cart.push({
            id: product.ID, // Store product.ID (from backend)
            name: product.name,
            price: product.price,
            quantity: quantity
        });
    }
    saveCart(cart);
    alert(`${product.name} added to cart!`);
}

export function updateCartItemQuantity(productId, quantity) {
    let cart = getCart();
    const itemIndex = cart.findIndex(item => item.id === productId);
    if (itemIndex > -1) {
        if (quantity > 0) {
            cart[itemIndex].quantity = quantity;
        } else {
            cart.splice(itemIndex, 1); // Remove if quantity is 0 or less
        }
        saveCart(cart);
    }
}

export function removeFromCart(productId) {
    let cart = getCart();
    cart = cart.filter(item => item.id !== productId);
    saveCart(cart);
}

export function clearCart() {
    localStorage.removeItem('cart');
    updateCartCount();
}

export function updateCartCount() {
    const cart = getCart();
    const cartCountElement = document.getElementById('cart-count');
    if (cartCountElement) {
        cartCountElement.textContent = cart.reduce((sum, item) => sum + item.quantity, 0);
    }
}

// Function to display messages
export function displayMessage(message, type = 'success', containerId = 'message-container') {
    const container = document.getElementById(containerId);
    if (!container) {
        // Create a default message container if one isn't specified or found
        const defaultContainer = document.createElement('div');
        defaultContainer.id = 'default-message-container';
        defaultContainer.style.position = 'fixed';
        defaultContainer.style.top = '20px';
        defaultContainer.style.left = '50%';
        defaultContainer.style.transform = 'translateX(-50%)';
        defaultContainer.style.zIndex = '1000';
        document.body.appendChild(defaultContainer);
        // container = defaultContainer; // This line has an issue, container is const
        // Correct approach:
        const messageContainer = document.getElementById('default-message-container') || defaultContainer;
        messageContainer.innerHTML = `<div class="alert alert-${type}" role="alert">${message}</div>`;
        setTimeout(() => { messageContainer.innerHTML = ''; }, 5000); // Clear after 5s
        return;
    }

    container.innerHTML = `<div class="alert alert-${type}" role="alert">${message}</div>`;
    setTimeout(() => { container.innerHTML = ''; }, 5000); // Clear after 5 seconds
}

// Debounce function
export function debounce(func, delay) {
    let timeout;
    return function(...args) {
        const context = this;
        clearTimeout(timeout);
        timeout = setTimeout(() => func.apply(context, args), delay);
    };
}