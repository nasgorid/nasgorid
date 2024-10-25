// Function to format numbers as Rupiah
function formatRupiah(number) {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(number);
}

// Function to fetch products from the backend
async function fetchProducts() {
    try {
        const response = await fetch('http://localhost:8081/products'); // Fetch the products using the GET route
        const products = await response.json(); // Parse the response as JSON

        // Find the specific product table container
        const productTable = document.getElementById('product-table');

        // Clear the existing rows (if necessary)
        document.querySelectorAll('.row.product').forEach(row => row.remove());

        // Loop through the products and create new rows
        products.forEach(product => {
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

            // Append the row to the table
            productTable.appendChild(row);
        });

        // Add event listeners for Edit and Delete buttons
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

    } catch (error) {
        console.error('Error fetching products:', error);
    }
}

// Example function to handle the edit action
function editProduct(productId) {
    // Redirect or perform other actions to edit the product with the specified ID
    console.log('Edit product with ID:', productId);
    window.location.href = `Edit_product.html?id=${productId}`;
}

// Example function to handle the delete action
// Example function to handle the delete action
async function deleteProduct(productId) {
    // Tampilkan pesan konfirmasi sebelum menghapus
    const isConfirmed = window.confirm("Apakah Anda yakin ingin menghapus produk ini?");
    
    if (!isConfirmed) {
        return; // Jika pengguna membatalkan, hentikan proses penghapusan
    }

    try {
        const response = await fetch(`http://localhost:8081/products/${productId}`, { method: 'DELETE' });
        if (response.ok) {
            alert("Produk berhasil dihapus."); // Pesan sukses
            fetchProducts(); // Refresh the product list
        } else {
            alert("Gagal menghapus produk."); // Pesan error jika gagal
        }
    } catch (error) {
        console.error('Error deleting product:', error);
        alert("Terjadi kesalahan saat menghapus produk."); // Pesan error jika ada error jaringan
    }
}


// Call the function to fetch and display products when the page loads
window.onload = fetchProducts;

document.getElementById('exportCsvBtn').addEventListener('click', function() {
    // Redirect to the CSV export endpoint
    window.location.href = "http://localhost:8081/products-export-csv";
});

document.getElementById('addProductBtn').addEventListener('click', function() {
    // Redirect to AddProduct.html
    window.location.href = "AddProduct.html";
});
