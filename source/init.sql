CREATE TABLE users
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    mobile_id   VARCHAR(255) UNIQUE NOT NULL,
    password    VARCHAR(255)        NOT NULL,
    use_type    ENUM ('PLACE') DEFAULT 'PLACE',
    create_date TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_date TIMESTAMP           NULL
);

CREATE TABLE products
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    category    VARCHAR(255),
    user_id     INT,
    name        VARCHAR(255),
    initial     VARCHAR(255),
    price       DECIMAL(10, 2),
    cost        DECIMAL(10, 2),
    description TEXT,
    barcode     VARCHAR(50),
    expiry_date TIMESTAMP NOT NULL,
    size        VARCHAR(50),
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_date TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    INDEX idx_products_initial (initial),
    INDEX idx_products_name (name)
);

CREATE TABLE auth_tokens
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    user_id         INT,
    jwt_token       TEXT,
    creation_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiration_time TIMESTAMP NULL,
    active          BOOLEAN   DEFAULT TRUE,
    FOREIGN KEY (user_id) REFERENCES users (id),
    INDEX idx_auth_tokens_user_id (user_id),
    INDEX idx_auth_tokens_jwt_token (jwt_token(255))
);

INSERT INTO users (mobile_id, password) VALUES ('01011111111', '$2a$10$y8k/LZCyzGRnWlFCB2DzOenf5cQWbsUQGyISzulWww.trbs4FwQeq');
INSERT INTO users (mobile_id, password) VALUES ('01022222222', '$2a$10$y8k/LZCyzGRnWlFCB2DzOenf5cQWbsUQGyISzulWww.trbs4FwQeq');

INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '아메리카노', 'ㅇㅁㄹㅋㄴ', 3000, 1500, '아메리카노 판매합니다.', '12345678', '2024-03-01 09:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '카페라떼', 'ㅋㅍㄹㄸ', 3500, 1800, '카페라떼 판매합니다.', '23456789', '2024-03-01 09:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '카페모카', 'ㅋㅍㅁㅋ', 3800, 2000, '카페모카 판매합니다.', '34567890', '2024-03-01 10:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '헤이즐넛라떼', 'ㅎㅇㅈㄴㄹㄸ', 4000, 2000, '헤이즐넛라떼 판매합니다.', '45678901', '2024-03-01 10:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '바닐라라떼', 'ㅂㄷㄴㄹㄸ', 4000, 2000, '바닐라라떼 판매합니다.', '56789012', '2024-03-01 11:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '카푸치노', 'ㅋㅍㅊㄴ', 3700, 1900, '카푸치노 판매합니다.', '67890123', '2024-03-01 11:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '모카라떼', 'ㅁㅋㄹㄸ', 3900, 2000, '모카라떼 판매합니다.', '78901234', '2024-03-01 12:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '콜드브루', 'ㅋㄷㅂㄹ', 4500, 2200, '콜드브루 판매합니다.', '89012345', '2024-03-01 12:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '아이스티', 'ㅇㅇㅅㅌ', 3200, 1600, '아이스티 판매합니다.', '90123456', '2024-03-01 13:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '스무디', 'ㅅㅁㄷ', 5000, 2500, '스무디 판매합니다.', '01234567', '2024-03-01 13:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '플레인요거트', 'ㅍㄹㅇㅇㄱㅌ', 5500, 2700, '플레인요거트 판매합니다.', '12345678', '2024-03-01 14:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '딸기요거트', 'ㄸㄱㅇㄱㅌ', 5800, 2800, '딸기요거트 판매합니다.', '23456789', '2024-03-01 14:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '딸기 요거트', 'ㄸㄱ ㅇㄱㅌ', 5500, 2700, '딸기 요거트 판매합니다.', '12345678', '2024-03-01 14:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '블루베리 요거트', 'ㅂㄹㅂㄹ ㅇㄱㅌ', 5800, 2800, '블루베리 요거트 판매합니다.', '23456789', '2024-03-01 14:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '치즈 케이크', 'ㅊㅈ ㅋㅇㅋ', 7000, 3500, '치즈 케이크 판매합니다.', '34567890', '2024-03-01 15:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '초코 브라우니', 'ㅊㅋ ㅂㄹㅇㄴ', 6000, 3000, '초코 브라우니 판매합니다.', '45678901', '2024-03-01 15:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '카라멜 마카롱', 'ㅋㄹㅁ ㅁㅋㄹ', 6500, 3200, '카라멜 마카롱 판매합니다.', '56789012', '2024-03-01 16:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '말차 빙수', 'ㅁㅊ ㅂㅅ', 7500, 3700, '말차 빙수 판매합니다.', '67890123', '2024-03-01 16:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '아이스크림', 'ㅇㅇㅅㅋㄹ', 4000, 2000, '아이스크림 판매합니다.', '78901234', '2024-03-01 17:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '딸기 쉐이크', 'ㄸㄱ ㅅㅇㅋ', 4800, 2400, '딸기 쉐이크 판매합니다.', '89012345', '2024-03-01 17:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '바나나 크림', 'ㅂㄴㄴ ㅋㄹ', 5500, 2700, '바나나 크림 판매합니다.', '90123456', '2024-03-01 18:00:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('payhere', 1, '망고 스무디', 'ㅁㄱ ㅅㅁㄷ', 6300, 3100, '망고 스무디 판매합니다.', '01234567', '2024-03-01 18:30:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '나이키 운동화', 'ㄴㅇㅋ ㅇㄷㅎ', 80000, 50000, '나이키 운동화 판매합니다.', '12345678', '2024-03-01 14:00:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '아디다스 운동화', 'ㅇㄷㄷㅅ ㅇㄷㅎ', 90000, 60000, '아디다스 운동화 판매합니다.', '23456789', '2024-03-01 14:30:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '지오다노 티셔츠', 'ㅈㅇㄷㄴ ㅌㅅㅊ', 35000, 25000, '지오다노 티셔츠 판매합니다.', '34567890', '2024-03-01 15:00:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '폴로 셔츠', 'ㅍㄹ ㅅㅊ', 45000, 30000, '폴로 셔츠 판매합니다.', '45678901', '2024-03-01 15:30:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '구찌 반지갑', 'ㄱㅉ ㅂㅈㄱ', 150000, 100000, '구찌 반지갑 판매합니다.', '56789012', '2024-03-01 16:00:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '루이비통 가방', 'ㄹㅇㅂㅌ ㄱㅂ', 300000, 200000, '루이비통 가방 판매합니다.', '67890123', '2024-03-01 16:30:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '샤넬 향수', 'ㅅㄴ ㅎㅅ', 250000, 150000, '샤넬 향수 판매합니다.', '78901234', '2024-03-01 17:00:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '에르메스 벨트', 'ㅇㄹㅁㅅ ㅂㅌ', 180000, 120000, '에르메스 벨트 판매합니다.', '89012345', '2024-03-01 17:30:00', 'large');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '디올 클러치백', 'ㄷㅇ ㅋㄹㅊㅂ', 220000, 180000, '디올 클러치백 판매합니다.', '90123456', '2024-03-01 18:00:00', 'small');
INSERT INTO products (category, user_id, name, initial, price, cost, description, barcode, expiry_date, size) VALUES ('fashion', 1, '프라다 선글라스', 'ㅍㄹㄷ ㅅㄱㄹㅅ', 200000, 160000, '프라다 선글라스 판매합니다.', '01234567', '2024-03-01 18:30:00', 'large');