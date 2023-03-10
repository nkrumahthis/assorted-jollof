INSERT OR REPLACE INTO users(id, name, email, password) VALUES
    (1, 'Adam Asante', 'aa@assortedjollof.com', "password"),
    (2, 'Bill Boakye', 'bb@assortedjollof.com', "password"),
    (3, 'Charles C', 'cc@assortedjollof.com', "password");

INSERT OR REPLACE INTO customers(id, name, phone, token) VALUES
    (1, 'Derek Darko', '+1234567890', "1234"),
    (2, 'Eric Enchil', '+9872340192', "5678"),
    (3, 'Fred Faraday', '+7232019923', "9012");
INSERT OR REPLACE INTO orders(id, packs, customer_id, location, status, created_at) VALUES
    (1, 1, 1, 'Ashaiman Junction', 'FULFILLED', NULL),
    (2, 2, 2, 'Burma Camp', 'FULFILLED', NULL),
    (3, 1, 3, 'Next to China Mall', 'CANCELLED', NULL),
    (4, 2, 1, 'Ashaiman Junction', 'CANCELLED', NULL),
    (5, 3, 2, 'Burma Camp', 'FULFILLED', NULL),
    (6, 1, 1, 'Ashaiman Junction', 'FULFILLED', NULL),
    (7, 1, 3, 'Next to China Mall', 'PENDING', NULL),
    (8, 2, 1, 'Ashaiman Junction', 'PENDING', NULL),
    (9, 2, 2, 'Ashaiman Junction', 'PENDING', NULL);

INSERT OR REPLACE INTO payments(id, amount, customer_id, order_id, created_at) VALUES
    (1, 45.00, 1, 1, NULL),
    (2, 90.00, 2, 2, NULL),
    (3, 135.00, 2, 5, NULL),
    (4, 45.00, 1, 6, NULL);

