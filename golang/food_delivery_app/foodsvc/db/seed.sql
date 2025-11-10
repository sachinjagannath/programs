-- Insert sample restaurants
INSERT INTO restaurants (name, location) VALUES
                                             ('Pizza Palace', 'Downtown'),
                                             ('Burger Barn', 'Uptown'),
                                             ('Sushi Central', 'Midtown');

-- Insert sample menu items
INSERT INTO menu_items (restaurant_id, name, price, available) VALUES
                                                                   (1, 'Margherita Pizza', 8.99, TRUE),
                                                                   (1, 'Pepperoni Pizza', 9.99, TRUE),
                                                                   (1, 'Veggie Pizza', 7.99, TRUE),
                                                                   (2, 'Classic Burger', 6.49, TRUE),
                                                                   (2, 'Cheeseburger', 6.99, TRUE),
                                                                   (2, 'Bacon Burger', 7.49, TRUE),
                                                                   (3, 'California Roll', 5.99, TRUE),
                                                                   (3, 'Spicy Tuna Roll', 6.49, TRUE),
                                                                   (3, 'Salmon Nigiri', 4.99, TRUE);

-- Insert sample orders (some pending, some completed)
INSERT INTO orders (menu_item_id, restaurant_id, customer_name, status) VALUES
                                                                            (1, 1, 'Alice', 'pending'),
                                                                            (2, 1, 'Bob', 'completed'),
                                                                            (4, 2, 'Charlie', 'pending'),
                                                                            (5, 2, 'David', 'completed'),
                                                                            (7, 3, 'Eve', 'pending'),
                                                                            (8, 3, 'Frank', 'completed');
