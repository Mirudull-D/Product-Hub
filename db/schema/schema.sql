CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       first_name TEXT NOT NULL,
                       last_name TEXT NOT NULL,
                       email TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL
);

CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT NOT NULL,
                          image TEXT NOT NULL,
                          quantity INT NOT NULL DEFAULT 0,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                        total DECIMAL(10, 2) NOT NULL,
                        status TEXT NOT NULL DEFAULT 'pending',
                        address TEXT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE orders_items (
                              id BIGSERIAL PRIMARY KEY,
                              order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
                              product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
                              quantity INT NOT NULL,
                              price DECIMAL(10, 2) NOT NULL
);