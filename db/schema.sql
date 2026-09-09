    CREATE TABLE IF NOT EXISTS venues (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        location TEXT NOT NULL,
        region TEXT NOT NULL,
        link TEXT
    );
    CREATE TABLE IF NOT EXISTS labels (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        location TEXT,
        region TEXT NOT NULL,
        genre TEXT NOT NULL,
        link TEXT
    );
    CREATE TABLE IF NOT EXISTS artists (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        dj_name TEXT NOT NULL,
        label TEXT,
        location TEXT,
        region TEXT,
        genre TEXT NOT NULL,
        link TEXT
    );
    CREATE TABLE IF NOT EXISTS genres (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS genre_links (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        genre_id INTEGER NOT NULL,
        history_link TEXT NOT NULL,
        FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
    );

