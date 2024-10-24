// Function to fetch products from the backend
async function fetchProducts() {
    try {
      const response = await fetch('http://localhost:8081/products'); // Fetch the products using the GET route
      const products = await response.json(); // Parse the response as JSON
  
      // Find the container where the product rows will be appended
      const productTable = document.querySelector('.container');
  
      // Clear the existing rows (if necessary)
      document.querySelectorAll('.row.product').forEach(row => row.remove());
  
      // Loop through the products and create new rows
      products.forEach(product => {
        const row = document.createElement('div');
        row.classList.add('row', 'product');
  
        row.innerHTML = `
          <div class="cell" data-title="Product">${product.name}</div>
          <div class="cell" data-title="Unit Price">$${product.price}</div>
          <div class="cell" data-title="Category">${product.category}</div>
          <div class="cell" data-title="Description">${product.description}</div>
          <div class="cell" data-title="Stock">${product.stock}</div>
        `;
  
        // Append the row to the table
        productTable.appendChild(row);
      });
    } catch (error) {
      console.error('Error fetching products:', error);
    }
  }
  
  // Call the function to fetch and display products when the page loads
  window.onload = fetchProducts;
  