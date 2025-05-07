// src/api/products.js
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export async function getAllProducts(page = 1, pageSize = 10) {
    const response = await fetch(`${API_BASE_URL}/products?page=${page}&pageSize=${pageSize}`);
    if (!response.ok) {
        console.error("Failed to fetch products:", response.status, await response.text());
        return { data: [], total: 0, page: 1, last_page: 1 }; // Return default on error
    }
    return response.json();
}

export async function getProductById(id) {
    const response = await fetch(`${API_BASE_URL}/products/${id}`);
    return response.json();
}

export async function createProduct(productData, token) {
    const response = await fetch(`${API_BASE_URL}/admin/products`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(productData),
    });
    return response.json();
}

export async function updateProduct(id, productData, token) {
    const response = await fetch(`${API_BASE_URL}/admin/products/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(productData),
    });
    return response.json();
}

export async function deleteProduct(id, token) {
    const response = await fetch(`${API_BASE_URL}/admin/products/${id}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    });
    // DELETE might not return a body, or an empty one if successful
    if (response.ok) return { success: true, status: response.status };
    return response.json(); // For error messages
}