CREATE TABLE restaurants (
                             id BIGINT AUTO_INCREMENT PRIMARY KEY,
                             name VARCHAR(255),
                             location VARCHAR(255)
);

CREATE TABLE menu_items (
                            id BIGINT AUTO_INCREMENT PRIMARY KEY,
                            restaurant_id BIGINT,
                            name VARCHAR(255),
                            price DECIMAL(10,2),
                            available BOOLEAN DEFAULT TRUE,
                            FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);

CREATE TABLE orders (
                        id BIGINT AUTO_INCREMENT PRIMARY KEY,
                        menu_item_id BIGINT,
                        restaurant_id BIGINT,
                        customer_name VARCHAR(255),
                        status VARCHAR(50),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (menu_item_id) REFERENCES menu_items(id),
                        FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);
