-- +goose Up
CREATE TABLE terms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    expression TEXT NOT NULL,
    reading TEXT NOT NULL,
    meaning TEXT NOT NULL
);

CREATE TABLE kanjis (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    character TEXT NOT NULL,
    meaning TEXT NOT NULL,
    on_yomi TEXT,
    kun_yomi TEXT,
    level TEXT
);

CREATE TABLE lists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE kanji_term (
    term_id INTEGER NOT NULL,
    kanji_id INTEGER NOT NULL,
    PRIMARY KEY (term_id, kanji_id),
    FOREIGN KEY (term_id) REFERENCES terms(id) ON DELETE CASCADE,
    FOREIGN KEY (kanji_id) REFERENCES kanji(id) ON DELETE CASCADE
);

CREATE TABLE category_term (
    term_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY (term_id, category_id),
    FOREIGN KEY (term_id) REFERENCES terms(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE list_term (
    term_id INTEGER NOT NULL,
    list_id INTEGER NOT NULL,
    PRIMARY KEY (term_id, list_id),
    FOREIGN KEY (term_id) REFERENCES terms(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE  
);

CREATE TABLE list_kanji (
    kanji_id INTEGER NOT NULL,
    list_id INTEGER NOT NULL,
    PRIMARY KEY (kanji_id, list_id),
    FOREIGN KEY (kanji_id) REFERENCES kanjis(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE  
);



-- +goose Down
DROP TABLE IF EXISTS list_kanji;
DROP TABLE IF EXISTS list_term;
DROP TABLE IF EXISTS category_term;
DROP TABLE IF EXISTS kanji_term;

DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS lists;
DROP TABLE IF EXISTS kanjis;
DROP TABLE IF EXISTS terms;
