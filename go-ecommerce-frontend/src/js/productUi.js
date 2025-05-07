// src/js/productUi.js
import { getAllProducts, getProductById } from '../api/products.js';
import { addToCart, displayMessage } from './utils.js';

export async function loadProducts(page = 1, pageSize = 8) {
    const productGrid = document.getElementById('product-grid');
    const paginationControls = document.getElementById('pagination-controls');
    if (!productGrid) return;

    productGrid.innerHTML = '<p>Loading products...</p>'; // Clear previous products

    try {
        const result = await getAllProducts(page, pageSize);
        // The backend product model has Name, Description, Price, Stock, CategoryID, Model (ID, CreatedAt etc)
        // The frontend expects products to be an array in result.data for pagination
        // Backend response structure: { data: [products], total: number, page: number, last_page: number }
        const products = result.data;
        const totalProducts = result.total;
        const currentPage = result.page;
        const lastPage = result.last_page;


        if (products && products.length > 0) {
            productGrid.innerHTML = ''; // Clear loading message
            products.forEach(product => {
                const productCard = document.createElement('div');
                productCard.className = 'product-card';
                productCard.innerHTML = `
                    <a href="/src/pages/product-detail.html?id=${product.ID}">
                        <img src="https://via.placeholder.com/200x150?text=${encodeURIComponent(product.name)}" alt="${product.name}">
                        <h3>${product.name}</h3>
                    </a>
                    <p class="price">$${product.price.toFixed(2)}</p>
                    <p>Stock: ${product.stock}</p>
                    <button class="add-to-cart-btn" data-product-id="${product.ID}">Add to Cart</button>
                `;
                productGrid.appendChild(productCard);

                // Add event listener for the new button
                productCard.querySelector('.add-to-cart-btn').addEventListener('click', () => {
                    // We need the full product object for addToCart, or at least ID, name, price
                    // The current 'product' in scope here has all details from the API
                    addToCart(product);
                });
            });

            // Pagination
            if (paginationControls) {
                paginationControls.innerHTML = ''; // Clear old controls
                if (lastPage > 1) {
                    const prevButton = document.createElement('button');
                    prevButton.textContent = 'Previous';
                    prevButton.disabled = currentPage === 1;
                    prevButton.addEventListener('click', () => loadProducts(currentPage - 1, pageSize));
                    paginationControls.appendChild(prevButton);

                    for (let i = 1; i <= lastPage; i++) {
                        const pageButton = document.createElement('button');
                        pageButton.textContent = i;
                        if (i === currentPage) pageButton.disabled = true;
                        pageButton.addEventListener('click', () => loadProducts(i, pageSize));
                        paginationControls.appendChild(pageButton);
                    }

                    const nextButton = document.createElement('button');
                    nextButton.textContent = 'Next';
                    nextButton.disabled = currentPage === lastPage;
                    nextButton.addEventListener('click', () => loadProducts(currentPage + 1, pageSize));
                    paginationControls.appendChild(nextButton);
                }
            }

        } else if (products && products.length === 0) {
            productGrid.innerHTML = '<p>No products found.</p>';
        } else { // Handle error case from API if result.data is undefined/null
             productGrid.innerHTML = `<p>Error loading products: ${result.error || 'Unknown error'}</p>`;
        }
    } catch (error) {
        console.error('Error fetching products:', error);
        productGrid.innerHTML = '<p>Could not load products. Please try again later.</p>';
    }
}


export async function initProductDetailPage() {
    const params = new URLSearchParams(window.location.search);
    const productId = params.get('id');
    const productDetailContainer = document.getElementById('product-detail-container'); // You'll need this div in product-detail.html

    if (!productId || !productDetailContainer) {
        if (productDetailContainer) productDetailContainer.innerHTML = "<p>Product ID not found.</p>";
        return;
    }

    try {
        const product = await getProductById(productId);
        if (product && product.ID) { // Check if product has ID (meaning it was found)
            productDetailContainer.innerHTML = `
                <h2>${product.name}</h2>
                <img src="https://via.placeholder.com/400x300?text=${encodeURIComponent(product.name)}" alt="${product.name}" style="max-width:100%;">
                <p>${product.description || 'No description available.'}</p>
                <p class="price">Price: $${product.price.toFixed(2)}</p>
                <p>Stock: ${product.stock}</p>
                <p>Category ID: ${product.category_id}</p> <!-- Assuming category_id is available -->
                <div class="form-group">
                    <label for="quantity">Quantity:</label>
                    <input type="number" id="quantity" value="1" min="1" max="${product.stock}">
                </div>
                <button id="add-to-cart-detail-btn">Add to Cart</button>
                <div id="message-container-detail"></div>
            `;
            document.getElementById('add-to-cart-detail-btn').addEventListener('click', () => {
                const quantity = parseInt(document.getElementById('quantity').value);
                if (quantity > 0 && quantity <= product.stock) {
                    addToCart(product, quantity);
                     displayMessage(`${product.name} (x${quantity}) added to cart!`, 'success', 'message-container-detail');
                } else {
                    displayMessage('Invalid quantity or out of stock.', 'danger', 'message-container-detail');
                }
            });
        } else {
            productDetailContainer.innerHTML = `<p>Product not found or error loading details: ${product.error || ''}</p>`;
        }
    } catch (error) {
        console.error('Error fetching product details:', error);
        productDetailContainer.innerHTML = '<p>Could not load product details.</p>';
    }
}