// src/api/orders.js
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function createOrder(orderData, token) {
    // Example orderData: { items: [{ product_id: 1, quantity: 2 }, ...] }
    const response = await fetch(`${API_BASE_URL}/orders`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(orderData),
    });
    return response.json();
}

export async function getUserOrders(token) {
    const response = await fetch(`${API_BASE_URL}/orders`, { // Assuming GET /orders for user's orders
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });
    return response.json();
}

// Add admin order functions if needed, e.g., getAllOrdersAdmin, updateOrderStatusAdmin