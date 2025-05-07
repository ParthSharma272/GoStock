// src/js/checkoutUi.js
import { getCart, clearCart, getToken, displayMessage } from './utils.js';
import { createOrder } from '../api/orders.js';

export function initCheckoutPage() {
    const token = getToken();
    const checkoutItemsContainer = document.getElementById('checkout-items');
    const checkoutTotalElement = document.getElementById('checkout-total');
    const orderForm = document.getElementById('order-form');
    const loginPrompt = document.getElementById('login-prompt');
    const shippingFormContainer = document.getElementById('shipping-details-form');

    if (!checkoutItemsContainer || !checkoutTotalElement || !orderForm || !loginPrompt || !shippingFormContainer) {
        console.error("Checkout page elements not found.");
        return;
    }

    if (!token) {
        shippingFormContainer.style.display = 'none';
        loginPrompt.style.display = 'block';
        checkoutItemsContainer.innerHTML = '<p>Please log in to view your cart summary for checkout.</p>';
        return;
    } else {
        shippingFormContainer.style.display = 'block';
        loginPrompt.style.display = 'none';
    }

    renderCheckoutSummary();

    function renderCheckoutSummary() {
        const cart = getCart();
        checkoutItemsContainer.innerHTML = '';
        let total = 0;

        if (cart.length === 0) {
            checkoutItemsContainer.innerHTML = '<p>Your cart is empty. <a href="/">Continue shopping</a>.</p>';
            checkoutTotalElement.textContent = '0.00';
            if (orderForm) orderForm.style.display = 'none'; // Hide form if cart is empty
            return;
        } else {
            if (orderForm) orderForm.style.display = 'block';
        }

        const ul = document.createElement('ul');
        ul.style.listStyleType = 'none';
        ul.style.paddingLeft = '0';

        cart.forEach(item => {
            const itemTotal = item.price * item.quantity;
            total += itemTotal;
            const li = document.createElement('li');
            li.textContent = `${item.name} (x${item.quantity}) - $${itemTotal.toFixed(2)}`;
            ul.appendChild(li);
        });
        checkoutItemsContainer.appendChild(ul);
        checkoutTotalElement.textContent = total.toFixed(2);
    }

    if (orderForm) {
        orderForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const placeOrderBtn = document.getElementById('place-order-btn');
            placeOrderBtn.disabled = true;
            placeOrderBtn.textContent = 'Placing Order...';

            const cart = getCart();
            if (cart.length === 0) {
                displayMessage('Your cart is empty.', 'danger');
                placeOrderBtn.disabled = false;
                placeOrderBtn.textContent = 'Place Order';
                return;
            }

            // Backend expects items: [{ product_id: X, quantity: Y }, ...]
            const orderItems = cart.map(item => ({
                product_id: item.id, // Ensure your cart items store product.ID as 'id'
                quantity: item.quantity,
            }));

            const shippingAddress = { // Example, your backend might want a different structure
                address: e.target.address.value,
                city: e.target.city.value,
                postal_code: e.target.postal_code.value,
            };

            const orderData = {
                items: orderItems,
                shipping_address: shippingAddress, // Send shipping address if your backend handles it
                // Payment details would go here if integrating a payment gateway
            };

            try {
                const result = await createOrder(orderData, token);
                // Assuming backend order creation responds with something like:
                // { order_id: 123, message: "Order created successfully", ... }
                // or { error: "Something went wrong" }
                if (result.order_id || (result.ID && result.status)) { // Check for a successful order indicator
                    displayMessage(`Order placed successfully! Order ID: ${result.order_id || result.ID}`, 'success');
                    clearCart();
                    setTimeout(() => {
                        window.location.href = '/src/pages/profile.html'; // Redirect to profile/order history
                    }, 2000);
                } else {
                    displayMessage(result.error || 'Failed to place order. Please try again.', 'danger');
                    placeOrderBtn.disabled = false;
                    placeOrderBtn.textContent = 'Place Order';
                }
            } catch (error) {
                console.error('Error placing order:', error);
                displayMessage('An unexpected error occurred. Please try again.', 'danger');
                placeOrderBtn.disabled = false;
                placeOrderBtn.textContent = 'Place Order';
            }
        });
    }
}