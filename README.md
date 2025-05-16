# GoStock - E-Commerce Platform

GoShop is a full-stack e-commerce application built with a Go backend and a vanilla JavaScript frontend (Vite).

## Core Features

*   **Backend (Go):** Product & Category (structure) Management, Inventory, Order Processing (basic), User Accounts (Admin/Customer roles), JWT Auth, RESTful API with Gin & GORM (PostgreSQL).
*   **Frontend (JS):** Product Browsing & Detail, Cart, Simplified Checkout, User Profile & Order History, Admin Product CRUD.

## Tech Stack

*   **Backend:** Go, Gin Gonic, GORM, PostgreSQL, JWT
*   **Frontend:** Vanilla JavaScript (ES6+), HTML5, CSS3, Vite

## Quick Start

### 1. Backend Setup (`go-ecommerce-backend/`)

1.  **Configure:** Copy `go-ecommerce-backend/.env.example` to `.env`. Update `DB_URL` for your PostgreSQL and set a strong `JWT_SECRET_KEY`. Ensure the database exists.
2.  **Project Path:** Replace `your_project_path` in `.go` files with your Go module path (from `go.mod`).
3.  **Dependencies:** `go mod tidy`
4.  **Run:** `go run cmd/api/main.go` (Defaults to `http://localhost:8080`)

### 2. Frontend Setup (`go-ecommerce-frontend/`)

1.  **Configure:** Ensure `go-ecommerce-frontend/.env` has `VITE_API_BASE_URL=http://localhost:8080/api/v1` (or your backend URL).
2.  **Dependencies:** `npm install`
3.  **Run:** `npm run dev` (Defaults to `http://localhost:3000`)

## Usage

*   **Browse & Shop:** Access the frontend URL (e.g., `http://localhost:3000`).
*   **Admin Panel:** Log in as an admin (register with 'admin' role or update DB). Access admin product management via the navbar link or directly (`/src/pages/admin/products.html`).

## Future Work

*   Full Category & Order Management
*   Payment Integration (Stripe/PayPal)
*   Enhanced UI/UX, Search, Filtering
*   Proper DB Migrations & Testing
