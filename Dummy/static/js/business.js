function showSales() {
    document.getElementById('salesTable').classList.remove('hidden');
    document.getElementById('expensesTable').classList.add('hidden');
    fetchSalesData();
}

function showExpenses() {
    document.getElementById('expensesTable').classList.remove('hidden');
    document.getElementById('salesTable').classList.add('hidden');
    fetchExpensesData();
}

async function fetchSalesData() {
    const response = await fetch('http://localhost:8081/transaksi');
    const salesData = await response.json();
    const salesTableBody = document.getElementById('salesTableBody');
    salesTableBody.innerHTML = '';
    salesData.forEach(sale => {
        const row = `<tr>
            <td>${new Date(sale.transactionDate).toLocaleDateString()}</td>
            <td>${sale.customer_name}</td>
            <td>${sale.total_amount}</td>
            <td>${sale.payment_method}</td>
        </tr>`;
        salesTableBody.innerHTML += row;
    });
}

async function fetchExpensesData() {
    const response = await fetch('http://localhost:8081/expense');
    const expensesData = await response.json();
    const expensesTableBody = document.getElementById('expensesTableBody');
    expensesTableBody.innerHTML = '';
    expensesData.forEach(expense => {
        const row = `<tr>
            <td>${new Date(expense.expense_date).toLocaleDateString()}</td>
            <td>${expense.expense_name}</td>
            <td>${expense.amount}</td>
            <td>${expense.category}</td>
        </tr>`;
        expensesTableBody.innerHTML += row;
    });
}

async function fetchCustomers() {
    const response = await fetch('http://localhost:8081/customers');
    const customersData = await response.json();
    const customersTableBody = document.getElementById('customersTableBody');
    customersTableBody.innerHTML = '';
    customersData.forEach(customer => {
        const row = `<tr>
            <td>${customer.name}</td>
            <td>${customer.email}</td>
            <td>${customer.phone}</td>
            <td>${customer.address}</td>
        </tr>`;
        customersTableBody.innerHTML += row;
    });
}

// Fetch initial data
showSales();