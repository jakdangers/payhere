CREATE TABLE users
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    mobile_id VARCHAR(255) UNIQUE NOT NULL,
    password      VARCHAR(255)        NOT NULL,
    use_type      ENUM ('PLACE') DEFAULT 'PLACE',
    create_date   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    update_date   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_date   TIMESTAMP           NULL
);

CREATE TABLE categories
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    parent_id   INT,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_date TIMESTAMP    NULL,
    FOREIGN KEY (parent_id) REFERENCES categories (id)
);

CREATE TABLE products
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    category_id INT,
    user_id     INT,
    name        VARCHAR(255),
    initial     VARCHAR(255),
    price       DECIMAL(10, 2),
    description TEXT,
    barcode     VARCHAR(50),
    expiry_date DATE,
    size        VARCHAR(50),
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_date TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE auth_tokens
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    user_id         INT,
    jwt_token       TEXT,
    creation_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiration_time TIMESTAMP NULL,
    active          BOOLEAN   DEFAULT TRUE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);