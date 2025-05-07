// src/js/adminProductUi.js
import { getToken, getUserRole, displayMessage } from './utils.js';
import { getAllProducts, getProductById, createProduct, updateProduct, deleteProduct } from '../api/products.js';

function checkAdminAccess(deniedElementId, contentElementIdOrFormId) {
    const token = getToken();
    const role = getUserRole();
    const deniedDiv = document.getElementById(deniedElementId);
    const contentDivOrForm = document.getElementById(contentElementIdOrFormId);

    if (!token || role !== 'admin') {
        if (deniedDiv) deniedDiv.style.display = 'block';
        if (contentDivOrForm) contentDivOrForm.style.display = 'none';
        // Optionally redirect: window.location.href = '/';
        return false;
    }
    if (deniedDiv) deniedDiv.style.display = 'none';
    if (contentDivOrForm) contentDivOrForm.style.display = 'block'; // Or 'flex' etc. if it's a form
    return true;
}


export async function initAdminProductsPage(currentPage = 1, pageSize = 10) {
    if (!checkAdminAccess('admin-access-denied', 'admin-content')) return;

    const productListContainer = document.getElementById('admin-product-list');
    const paginationControls = document.getElementById('admin-pagination-controls');
    if (!productListContainer) return;

    productListContainer.innerHTML = '<p>Loading products...</p>';
    const token = getToken();

    try {
        const result = await getAllProducts(currentPage, pageSize); // Using the public one
        const products = result.data;
        const totalProducts = result.total;
        const page = result.page;
        const lastPage = result.last_page;

        if (products && products.length > 0) {
            const table = document.createElement('table');
            table.innerHTML = `
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Stock</th>
                        <th>Category ID</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody></tbody>
            `;
            const tbody = table.querySelector('tbody');
            products.forEach(product => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${product.ID}</td>
                    <td>${product.name}</td>
                    <td>$${product.price.toFixed(2)}</td>
                    <td>${product.stock}</td>
                    <td>${product.category_id}</td>
                    <td>
                        <a href="/src/pages/admin/edit-product.html?id=${product.ID}" class="button-secondary" style="margin-right: 5px;">Edit</a>
                        <button class="delete-product-btn button-danger" data-id="${product.ID}">Delete</button>
                    </td>
                `;
                tbody.appendChild(tr);
            });
            productListContainer.innerHTML = '';
            productListContainer.appendChild(table);

            // Event listeners for delete buttons
            productListContainer.querySelectorAll('.delete-product-btn').forEach(button => {
                button.addEventListener('click', async (e) => {
                    const productId = e.target.dataset.id;
                    if (confirm(`Are you sure you want to delete product ID ${productId}?`)) {
                        const deleteResult = await deleteProduct(productId, token);
                        if (deleteResult.success || deleteResult.status === 204 || deleteResult.status === 200) {
                            displayMessage(`Product ${productId} deleted successfully.`, 'success');
                            initAdminProductsPage(page, pageSize); // Refresh list
                        } else {
                            displayMessage(deleteResult.error || 'Failed to delete product.', 'danger');
                        }
                    }
                });
            });

             // Pagination for Admin
            if (paginationControls) {
                paginationControls.innerHTML = ''; // Clear old controls
                if (lastPage > 1) {
                    const prevButton = document.createElement('button');
                    prevButton.textContent = 'Previous';
                    prevButton.disabled = page === 1;
                    prevButton.addEventListener('click', () => initAdminProductsPage(page - 1, pageSize));
                    paginationControls.appendChild(prevButton);

                    for (let i = 1; i <= lastPage; i++) {
                        const pageButton = document.createElement('button');
                        pageButton.textContent = i;
                        if (i === page) pageButton.disabled = true;
                        pageButton.addEventListener('click', () => initAdminProductsPage(i, pageSize));
                        paginationControls.appendChild(pageButton);
                    }

                    const nextButton = document.createElement('button');
                    nextButton.textContent = 'Next';
                    nextButton.disabled = page === lastPage;
                    nextButton.addEventListener('click', () => initAdminProductsPage(page + 1, pageSize));
                    paginationControls.appendChild(nextButton);
                }
            }


        } else if (products && products.length === 0) {
            productListContainer.innerHTML = '<p>No products found.</p>';
        } else {
            productListContainer.innerHTML = `<p>Error loading products: ${result.error || 'Unknown error'}</p>`;
        }
    } catch (error) {
        console.error('Error fetching admin products:', error);
        productListContainer.innerHTML = '<p>Could not load products for admin.</p>';
    }
}


export function initAddProductPage() {
    const form = document.getElementById('add-product-form');
    if (!checkAdminAccess('admin-access-denied', 'add-product-form') || !form) return;

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const token = getToken();
        const productData = {
            name: e.target.name.value,
            description: e.target.description.value,
            price: parseFloat(e.target.price.value),
            stock: parseInt(e.target.stock.value),
            category_id: parseInt(e.target.category_id.value),
        };

        const result = await createProduct(productData, token);
        if (result.ID) { // Backend create product returns the product object with ID
            displayMessage(`Product "${result.name}" created successfully!`, 'success');
            form.reset();
            setTimeout(() => { window.location.href = '/src/pages/admin/products.html'; }, 1500);
        } else {
            displayMessage(result.error || 'Failed to create product.', 'danger');
        }
    });
}

export async function initEditProductPage() {
    const form = document.getElementById('edit-product-form');
    const loadingMsg = document.getElementById('loading-product-data');

    if (!checkAdminAccess('admin-access-denied', 'edit-product-form') || !form) return;


    const params = new URLSearchParams(window.location.search);
    const productId = params.get('id');

    if (!productId) {
        displayMessage('No product ID provided for editing.', 'danger');
        form.style.display = 'none';
        return;
    }

    if (loadingMsg) loadingMsg.style.display = 'block';
    form.style.display = 'none'; // Hide form while loading

    const token = getToken();

    try {
        const product = await getProductById(productId); // Public API to get product details
        if (product && product.ID) {
            form.elements.productId.value = product.ID;
            form.elements.name.value = product.name;
            form.elements.description.value = product.description || '';
            form.elements.price.value = product.price.toFixed(2);
            form.elements.stock.value = product.stock;
            form.elements.category_id.value = product.category_id;
            if (loadingMsg) loadingMsg.style.display = 'none';
            form.style.display = 'block'; // Show form once populated
        } else {
            displayMessage(product.error || 'Product not found or error loading data.', 'danger');
            if (loadingMsg) loadingMsg.style.display = 'none';
            return;
        }
    } catch (error) {
        displayMessage('Error fetching product data for editing.', 'danger');
        if (loadingMsg) loadingMsg.style.display = 'none';
        return;
    }


    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const productData = {
            name: e.target.name.value,
            description: e.target.description.value,
            price: parseFloat(e.target.price.value),
            stock: parseInt(e.target.stock.value),
            category_id: parseInt(e.target.category_id.value),
        };

        // Backend expects pointers for optional fields, so only send what's changed or send all.
        // For simplicity here, sending all. Your backend `UpdateProductRequest` uses pointers.
        // The API product.js needs to be smart about sending only non-empty fields or the backend
        // should handle `nil` for unchanged values if the DTO fields are pointers.
        // Let's assume the backend service layer handles `*string`, `*float64` correctly,
        // and the API layer `updateProduct` passes the struct correctly.

        const result = await updateProduct(productId, productData, token);
        if (result.ID) { // Backend update product returns the updated product object
            displayMessage(`Product "${result.name}" updated successfully!`, 'success');
            setTimeout(() => { window.location.href = '/src/pages/admin/products.html'; }, 1500);
        } else {
            displayMessage(result.error || 'Failed to update product.', 'danger');
        }
    });
}