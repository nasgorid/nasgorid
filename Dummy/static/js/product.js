// Function to format numbers as Rupiah
function formatRupiah(number) {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(number);
}

// Array untuk menyimpan data produk yang diambil dari backend
let products = [];

// Function to fetch products from the backend
async function fetchProducts() {
    try {
        const response = await fetch('http://localhost:8081/products'); // Fetch the products using the GET route
        products = await response.json(); // Parse the response as JSON

        // Render product table dengan data yang didapat
        renderProductTable(products);

        // Add event listeners for Edit and Delete buttons
        addEventListenersToButtons();

    } catch (error) {
        console.error('Error fetching products:', error);
    }
}

// Function to render product table
function renderProductTable(filteredProducts) {
    const productTable = document.getElementById('product-table');

    // Clear existing rows
    document.querySelectorAll('.row.product').forEach(row => row.remove());

    // Loop through the filtered products and create new rows
    filteredProducts.forEach(product => {
        const row = document.createElement('div');
        row.classList.add('row', 'product');

        row.innerHTML = `
            <div class="cell" data-title="Product">${product.name}</div>
            <div class="cell" data-title="Unit Price">${formatRupiah(product.price)}</div>
            <div class="cell" data-title="Category">${product.category}</div>
            <div class="cell" data-title="Description">${product.description}</div>
            <div class="cell" data-title="Stock">${product.stock}</div>
            <div class="cell">
                <button type="button" class="btn btn-success" data-id="${product.id}">Edit</button>
                <button type="button" class="btn btn-danger" data-id="${product.id}">Delete</button>
            </div>
        `;


        
        productTable.appendChild(row);
    });

    // Add event listeners to the new buttons
    addEventListenersToButtons();
}

// Function to add event listeners to Edit and Delete buttons
function addEventListenersToButtons() {
    document.querySelectorAll('.btn-success').forEach(button => {
        button.addEventListener('click', () => {
            const productId = button.getAttribute('data-id');
            editProduct(productId); // Call the edit function with the product ID
        });
    });

    document.querySelectorAll('.btn-danger').forEach(button => {
        button.addEventListener('click', () => {
            const productId = button.getAttribute('data-id');
            deleteProduct(productId); // Call the delete function with the product ID
        });
    });
}

// Function to handle search functionality
function searchProducts() {
    const searchQuery = document.getElementById('search-bar').value.toLowerCase();
    const filteredProducts = products.filter(product =>
        product.name.toLowerCase().includes(searchQuery) ||
        product.category.toLowerCase().includes(searchQuery) ||
        product.description.toLowerCase().includes(searchQuery)
    );

    renderProductTable(filteredProducts);
}

// Function to handle sorting functionality
function sortProducts() {
    const sortOption = document.getElementById('sort-options').value;

    const sortedProducts = [...products].sort((a, b) => {
        if (sortOption === "name") {
            return a.name.localeCompare(b.name);
        } else if (sortOption === "price") {
            return a.price - b.price;
        } else if (sortOption === "category") {
            return a.category.localeCompare(b.category);
        } else if (sortOption === "stock") {
            return a.stock - b.stock;
        }
    });

    renderProductTable(sortedProducts);
}

// Example function to handle the edit action
function editProduct(productId) {

    console.log('Edit product with ID:', productId);
    window.location.href = `Edit_product.html?id=${productId}`;
}


// Example function to handle the delete action
async function deleteProduct(productId) {

    const isConfirmed = window.confirm("Apakah Anda yakin ingin menghapus produk ini?");
    
    if (!isConfirmed) {
        return;
    }

    try {
        const response = await fetch(`http://localhost:8081/products/${productId}`, { method: 'DELETE' });
        if (response.ok) {
            alert("Produk berhasil dihapus.");
            fetchProducts(); // Refresh the product list
        } else {
            alert("Gagal menghapus produk.");
        }
    } catch (error) {
        console.error('Error deleting product:', error);
        alert("Terjadi kesalahan saat menghapus produk.");
    }
}

// Event listener for search and sort options
document.getElementById('search-bar').addEventListener('input', searchProducts);
document.getElementById('sort-options').addEventListener('change', sortProducts);

// Call the function to fetch and display products when the page loads
window.onload = fetchProducts;

// Event listeners for adding new product and exporting to CSV
document.getElementById('exportCsvBtn').addEventListener('click', function() {

    window.location.href = "http://localhost:8081/products-export-csv";
});

document.getElementById('addProductBtn').addEventListener('click', function() {

    window.location.href = "AddProduct.html";
});
