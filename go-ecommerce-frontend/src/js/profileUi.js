// src/js/profileUi.js
import { getToken } from './utils.js';
import { getUserProfile } from '../api/auth.js';
import { getUserOrders } from '../api/orders.js'; // Assuming this API endpoint exists

export async function initProfilePage() {
    const token = getToken();
    const profileDetailsContainer = document.getElementById('profile-details');
    const orderHistoryContainer = document.getElementById('order-history');
    const loginPrompt = document.getElementById('login-prompt-profile');

    if (!profileDetailsContainer || !orderHistoryContainer || !loginPrompt) {
        console.error("Profile page elements not found.");
        return;
    }

    if (!token) {
        profileDetailsContainer.style.display = 'none';
        orderHistoryContainer.style.display = 'none';
        loginPrompt.style.display = 'block';
        return;
    } else {
        loginPrompt.style.display = 'none';
    }

    // Load Profile Details
    try {
        const profile = await getUserProfile(token);
        if (profile && profile.id) { // Backend sends 'id' in /me response
            profileDetailsContainer.innerHTML = `
                <p><strong>Name:</strong> ${profile.first_name} ${profile.last_name}</p>
                <p><strong>Email:</strong> ${profile.email}</p>
                <p><strong>Role:</strong> ${profile.role}</p>
                <p><strong>Joined:</strong> ${new Date(profile.created_at).toLocaleDateString()}</p>
            `;
        } else {
            profileDetailsContainer.innerHTML = `<p>Could not load profile: ${profile.error || 'User not found or error.'}</p>`;
        }
    } catch (error) {
        console.error('Error fetching profile:', error);
        profileDetailsContainer.innerHTML = '<p>Error loading profile details.</p>';
    }

    // Load Order History
    // Assumes backend GET /api/v1/orders (when authenticated) returns an array of orders for the user.
    // Each order object might look like: { ID: 1, UserID: 1, TotalAmount: 50.00, Status: "pending", CreatedAt: "...", OrderItems: [...] }
    try {
        const ordersResult = await getUserOrders(token); // This should return an array directly or an object like { data: [] }
        let orders = [];
        if (Array.isArray(ordersResult)) {
            orders = ordersResult;
        } else if (ordersResult && Array.isArray(ordersResult.data)) {
            orders = ordersResult.data;
        } else if (ordersResult && ordersResult.error) {
            orderHistoryContainer.innerHTML = `<p>Could not load order history: ${ordersResult.error}</p>`;
            return;
        }


        if (orders.length > 0) {
            const table = document.createElement('table');
            table.innerHTML = `
                <thead>
                    <tr>
                        <th>Order ID</th>
                        <th>Date</th>
                        <th>Total</th>
                        <th>Status</th>
                        <th>Items</th>
                    </tr>
                </thead>
                <tbody></tbody>
            `;
            const tbody = table.querySelector('tbody');
            orders.forEach(order => {
                const tr = document.createElement('tr');
                // Assuming order items are nested or you make another call for details
                let itemsSummary = 'Details not shown';
                if (order.OrderItems && order.OrderItems.length > 0) {
                     itemsSummary = order.OrderItems.map(item => `${item.product_name || `Product ID ${item.product_id}`} (x${item.quantity})`).join(', ');
                } else if (order.items && order.items.length > 0) { // If backend sends 'items' directly
                     itemsSummary = order.items.map(item => `${item.product_name || `Product ID ${item.product_id}`} (x${item.quantity})`).join(', ');
                }


                tr.innerHTML = `
                    <td>#${order.ID}</td>
                    <td>${new Date(order.CreatedAt || order.created_at).toLocaleDateString()}</td>
                    <td>$${parseFloat(order.TotalAmount || order.total_amount || 0).toFixed(2)}</td>
                    <td>${order.Status || order.status}</td>
                    <td>${itemsSummary}</td>
                `;
                tbody.appendChild(tr);
            });
            orderHistoryContainer.innerHTML = ''; // Clear loading message
            orderHistoryContainer.appendChild(table);
        } else {
            orderHistoryContainer.innerHTML = '<p>You have no orders yet.</p>';
        }
    } catch (error) {
        console.error('Error fetching order history:', error);
        orderHistoryContainer.innerHTML = '<p>Error loading order history.</p>';
    }
}