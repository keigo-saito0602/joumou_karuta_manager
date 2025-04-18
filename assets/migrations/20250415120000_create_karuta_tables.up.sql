-- cards テーブル
CREATE TABLE cards (
  id SERIAL PRIMARY KEY,
  alphabet VARCHAR(255) NOT NULL,
  comment TEXT NOT NULL,
  initial VARCHAR(10) NOT NULL,
  picture_card VARCHAR(255) NOT NULL,
  text VARCHAR(255) NOT NULL,
  text_card VARCHAR(255) NOT NULL,
  yomi VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- event_results テーブル
CREATE TABLE event_results (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  feeling TEXT,
  result INTEGER NOT NULL,
  sub_result INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- results テーブル
CREATE TABLE results (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  feeling TEXT,
  result INTEGER NOT NULL,
  sub_result INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);