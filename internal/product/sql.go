package product

const findProductByIDQuery = "SELECT id, create_date, update_date, delete_date, user_id, initial, category, price, cost, name, description, barcode, expiry_date, size FROM `products` WHERE id = ?"

const updateProductQuery = "UPDATE `products` SET initial = ?, category = ?, price = ?, cost = ?, name = ?, description = ?, barcode = ?, expiry_date = ?, size = ? WHERE id = ?"

const deleteProductQuery = "UPDATE `products` SET delete_date = ? WHERE id = ?"
