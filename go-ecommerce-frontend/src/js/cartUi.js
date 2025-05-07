// src/js/cartUi.js (Simplified)
import { getCart, updateCartItemQuantity, removeFromCart, clearCart, displayMessage } from './utils.js';
// import { createOrder } from '../api/orders.js'; // For checkout

export function initCartPage() {
    const cartItemsContainer = document.getElementById('cart-items');
    const cartTotalElement = document.getElementById('cart-total');
    const checkoutButton = document.getElementById('checkout-btn'); // Add to cart.html
    const clearCartButton = document.getElementById('clear-cart-btn'); // Add to cart.html

    if (!cartItemsContainer) return;

    renderCart();

    function renderCart() {
        const cart = getCart();
        cartItemsContainer.innerHTML = ''; // Clear previous items
        let total = 0;

        if (cart.length === 0) {
            cartItemsContainer.innerHTML = '<p>Your cart is empty.</p>';
            if (cartTotalElement) cartTotalElement.textContent = '0.00';
            if (checkoutButton) checkoutButton.disabled = true;
            return;
        }

        const table = document.createElement('table');
        table.innerHTML = `
            <thead>
                <tr>
                    <th>Product</th>
                    <th>Price</th>
                    <th>Quantity</th>
                    <th>Subtotal</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody></tbody>
        `;
        const tbody = table.querySelector('tbody');

        cart.forEach(item => {
            const itemTotal = item.price * item.quantity;
            total += itemTotal;
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>${item.name}</td>
                <td>$${item.price.toFixed(2)}</td>
                <td>
                    <input type="number" value="${item.quantity}" min="1" class="cart-item-quantity" data-id="${item.id}" style="width: 60px;">
                </td>
                <td>$${itemTotal.toFixed(2)}</td>
                <td><button class="remove-from-cart-btn button-danger" data-id="${item.id}">Remove</button></td>
            `;
            tbody.appendChild(tr);
        });
        cartItemsContainer.appendChild(table);

        if (cartTotalElement) cartTotalElement.textContent = total.toFixed(2);
        if (checkoutButton) checkoutButton.disabled = false;

        // Add event listeners for quantity changes and removal
        cartItemsContainer.querySelectorAll('.cart-item-quantity').forEach(input => {
            input.addEventListener('change', (e) => {
                const productId = parseInt(e.target.dataset.id);
                const quantity = parseInt(e.target.value);
                updateCartItemQuantity(productId, quantity);
                renderCart(); // Re-render to update totals and potentially remove item if quantity is 0
            });
        });

        cartItemsContainer.querySelectorAll('.remove-from-cart-btn').forEach(button => {
            button.addEventListener('click', (e) => {
                const productId = parseInt(e.target.dataset.id);
                removeFromCart(productId);
                renderCart();
            });
        });
    }

    if (checkoutButton) {
        checkoutButton.addEventListener('click', () => {
            // Redirect to a dedicated checkout page or handle inline
            window.location.href = '/src/pages/checkout.html';
        });
    }

    if (clearCartButton) {
        clearCartButton.addEventListener('click', () => {
            if (confirm('Are you sure you want to clear your cart?')) {
                clearCart();
                renderCart();
                displayMessage('Cart cleared!', 'success');
            }
        });
    }
}