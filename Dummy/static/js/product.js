// Function to format numbers as Rupiah
function formatRupiah(number) {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(number);
}

// Array untuk menyimpan data produk yang diambil dari backend
let products = [];
let currentPage = 1;
const itemsPerPage = 8; // Number of products per page

// Function to fetch products from the backend
async function fetchProducts() {
    try {
        const response = await fetch('http://localhost:8081/products'); // Fetch the products using the GET route
        products = await response.json(); // Parse the response as JSON

        // Render the product table with initial page
        renderProductTable(products, currentPage);
        setupPagination(products);

    } catch (error) {
        console.error('Error fetching products:', error);
    }
}

// Function to render product table
function renderProductTable(productsArray, page) {
    const productTable = document.getElementById('product-table');

    // Clear existing rows
    document.querySelectorAll('.row.product').forEach(row => row.remove());

    // Calculate start and end index for the current page
    const startIndex = (page - 1) * itemsPerPage;
    const endIndex = page * itemsPerPage;
    const paginatedProducts = productsArray.slice(startIndex, endIndex);

    // Loop through the paginated products and create new rows
    paginatedProducts.forEach(product => {
        const row = document.createElement('div');
        row.classList.add('row', 'product');

        row.innerHTML = `
            <div class="cell" data-title="Product">${product.name}</div>
            <div class="cell" data-title="Unit Price">${formatRupiah(product.price)}</div>
            <div class="cell" data-title="Category">${product.category}</div>
            <div class="cell" data-title="Description">${product.description}</div>
            <div class="cell" data-title="Stock">${product.stock}</div>
            <div class="cell">
                <button type="button" class="btn btn-edit" data-id="${product.id}">
                    <i class="fas fa-pencil-alt"></i> <!-- Ikon pensil untuk edit -->
                </button>
                <button type="button" class="btn btn-delete" data-id="${product.id}">
                    <i class="fas fa-trash-alt"></i> <!-- Ikon tempat sampah untuk delete -->
                </button>
            </div>
        `;


        
        productTable.appendChild(row);
    });
}

// Event delegation for edit and delete buttons
document.getElementById('product-table').addEventListener('click', function(event) {
    const target = event.target;
    const productId = target.closest('button')?.getAttribute('data-id'); // Mengambil ID produk dari tombol terdekat

    if (target.closest('.btn-edit')) {
        editProduct(productId); // Panggil fungsi edit jika tombol edit diklik
    } else if (target.closest('.btn-delete')) {
        deleteProduct(productId); // Panggil fungsi delete jika tombol delete diklik
    }
});

// Function to setup pagination
function setupPagination(productsArray) {
    const paginationElement = document.getElementById('pagination');
    paginationElement.innerHTML = ''; // Clear existing pagination links

    const totalPages = Math.ceil(productsArray.length / itemsPerPage); // Calculate total pages

    // Create "Previous" button
    const prevButton = document.createElement('a');
    prevButton.href = "#";
    prevButton.innerHTML = "&laquo;";
    prevButton.addEventListener('click', () => {
        if (currentPage > 1) {
            currentPage--;
            renderProductTable(products, currentPage);
            updatePaginationLinks();
        }
    });
    paginationElement.appendChild(prevButton);

    // Create page number links
    for (let i = 1; i <= totalPages; i++) {
        const pageLink = document.createElement('a');
        pageLink.href = "#";
        pageLink.innerText = i;
        if (i === currentPage) {
            pageLink.classList.add('active');
        }
        pageLink.addEventListener('click', (event) => {
            currentPage = i;
            renderProductTable(products, currentPage);
            updatePaginationLinks();
        });
        paginationElement.appendChild(pageLink);
    }

    // Create "Next" button
    const nextButton = document.createElement('a');
    nextButton.href = "#";
    nextButton.innerHTML = "&raquo;";
    nextButton.addEventListener('click', () => {
        if (currentPage < totalPages) {
            currentPage++;
            renderProductTable(products, currentPage);
            updatePaginationLinks();
        }
    });
    paginationElement.appendChild(nextButton);
}

// Function to update the active page link in the pagination
function updatePaginationLinks() {
    const paginationLinks = document.querySelectorAll('#pagination a');
    paginationLinks.forEach(link => link.classList.remove('active'));

    // Highlight the current page link
    paginationLinks[currentPage].classList.add('active');
}

// Add event listeners to Edit and Delete buttons
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

// Function to edit a product
function editProduct(productId) {
    // Redirect to edit page, passing the product ID as a query parameter
    window.location.href = `EditProduct.html?id=${productId}`;
}


// Function to delete a product
async function deleteProduct(productId) {
    if (confirm("Are you sure you want to delete this product?")) {
        try {
            const response = await fetch(`http://localhost:8081/products/${productId}`, {
                method: 'DELETE',
            });

            if (response.ok) {
                alert("Product deleted successfully!");
                fetchProducts(); // Reload the product list after deletion
            } else {
                alert("Failed to delete the product.");
            }
        } catch (error) {
            console.error('Error deleting product:', error);
        }
    }
}

// Search Functionality
function searchProducts() {
    const searchQuery = document.getElementById('search-bar').value.toLowerCase();
    const filteredProducts = products.filter(product => 
        product.name.toLowerCase().includes(searchQuery) || 
        product.category.toLowerCase().includes(searchQuery) || 
        product.description.toLowerCase().includes(searchQuery)
    );
    renderProductTable(filteredProducts, currentPage);
    setupPagination(filteredProducts); // Update pagination based on filtered products
}

// Sort Functionality
function sortProducts() {
    const sortOption = document.getElementById('sort-options').value;
    let sortedProducts = [...products];


    
    if (sortOption === "name") {
        sortedProducts.sort((a, b) => a.name.localeCompare(b.name));
    } else if (sortOption === "price") {
        sortedProducts.sort((a, b) => a.price - b.price);
    } else if (sortOption === "category") {
        sortedProducts.sort((a, b) => a.category.localeCompare(b.category));
    } else if (sortOption === "stock") {
        sortedProducts.sort((a, b) => a.stock - b.stock);
    }

    renderProductTable(sortedProducts, currentPage);
    setupPagination(sortedProducts); // Update pagination based on sorted products
}

// Add event listeners for search and sort
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
