package product

const createProductQuery = "INSERT INTO products (user_id, initial, category, price, cost, name, description, barcode, expiry_date, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

const findProductByIDQuery = `
    SELECT 
        id,
        create_date,
        update_date,
        delete_date,
        user_id,
        initial, 
        category, 
        price, 
        cost,
        name,
        description, 
        barcode,
        expiry_date,
        size
    FROM 
        products 
    WHERE 
        delete_date IS NULL 
        AND id = ?
`

const updateProductQuery = `UPDATE products SET initial = ?, category = ?, price = ?, cost = ?, name = ?, description = ?, barcode = ?, expiry_date = ?, size = ? WHERE id = ?`

const deleteProductQuery = `UPDATE products SET delete_date = ? WHERE id = ?`

const listProductsQuery = `
	SELECT 
		id, 
		create_date, 
		update_date, 
		delete_date, 
		user_id, 
		initial, 
		category, 
		price, 
		cost, 
		name, 
		description, 
		barcode, 
		expiry_date, 
		size 
	FROM 
		products 
	WHERE 
		user_id = ? 
		AND delete_date IS NULL
		%s %s %s
	LIMIT 10
`
