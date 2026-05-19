-- ─── Users ────────────────────────────────────────────────────────────────────
CREATE TABLE users (
    id            TEXT        PRIMARY KEY,                          -- UUID (генерується в Go)
    email         TEXT        NOT NULL UNIQUE,
    password_hash TEXT        NOT NULL,
    name          TEXT        NOT NULL DEFAULT '',
    role          TEXT        NOT NULL DEFAULT 'user'
                              CHECK (role IN ('user', 'admin')),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── Categories ───────────────────────────────────────────────────────────────
CREATE TABLE categories (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE
);

-- ─── Plants ───────────────────────────────────────────────────────────────────
CREATE TABLE plants (
    id          SERIAL      PRIMARY KEY,
    name        TEXT        NOT NULL,
    slug        TEXT        NOT NULL UNIQUE,
    description TEXT        NOT NULL DEFAULT '',
    price       INTEGER     NOT NULL CHECK (price >= 0),   -- в центах (€10.00 = 1000)
    images      JSONB       NOT NULL DEFAULT '[]',          -- масив рядків-шляхів
    category_id INTEGER     REFERENCES categories(id) ON DELETE SET NULL,
    stock       INTEGER     NOT NULL DEFAULT 0 CHECK (stock >= 0),
    featured    BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_plants_category_id ON plants(category_id);
CREATE INDEX idx_plants_featured    ON plants(featured);
CREATE INDEX idx_plants_slug        ON plants(slug);

-- ─── Orders ───────────────────────────────────────────────────────────────────
CREATE TABLE orders (
    id               SERIAL      PRIMARY KEY,
    user_id          TEXT        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status           TEXT        NOT NULL DEFAULT 'pending'
                                 CHECK (status IN ('pending', 'paid', 'shipped', 'delivered')),
    total_price      INTEGER     NOT NULL CHECK (total_price >= 0),  -- в центах
    shipping_address JSONB       NOT NULL DEFAULT '{}',
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status  ON orders(status);

-- ─── Order Items ──────────────────────────────────────────────────────────────
CREATE TABLE order_items (
    id                SERIAL  PRIMARY KEY,
    order_id          INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    plant_id          INTEGER NOT NULL REFERENCES plants(id) ON DELETE RESTRICT,
    quantity          INTEGER NOT NULL DEFAULT 1 CHECK (quantity > 0),
    price_at_purchase INTEGER NOT NULL CHECK (price_at_purchase >= 0)  -- в центах
);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);

-- ─── Reviews ──────────────────────────────────────────────────────────────────
CREATE TABLE reviews (
    id         SERIAL      PRIMARY KEY,
    user_id    TEXT        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    plant_id   INTEGER     NOT NULL REFERENCES plants(id) ON DELETE CASCADE,
    rating     INTEGER     NOT NULL CHECK (rating BETWEEN 1 AND 5),
    text       TEXT        NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reviews_plant_id ON reviews(plant_id);
CREATE INDEX idx_reviews_user_id  ON reviews(user_id);

-- ─── Blog Posts ───────────────────────────────────────────────────────────────
CREATE TABLE blog_posts (
    id          SERIAL      PRIMARY KEY,
    title       TEXT        NOT NULL,
    slug        TEXT        NOT NULL UNIQUE,
    content     TEXT        NOT NULL DEFAULT '',  -- HTML з Tiptap
    excerpt     TEXT        NOT NULL DEFAULT '',
    cover_image TEXT,
    published   BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_blog_posts_published ON blog_posts(published);
CREATE INDEX idx_blog_posts_slug      ON blog_posts(slug);
