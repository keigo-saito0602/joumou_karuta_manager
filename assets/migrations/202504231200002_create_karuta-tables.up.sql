-- cards テーブル
CREATE TABLE cards (
  id BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'カードID',
  alphabet VARCHAR(255) NOT NULL COMMENT 'アルファベット（例：A〜Z）',
  comment TEXT NOT NULL COMMENT '補足コメント',
  initial VARCHAR(10) NOT NULL COMMENT '頭文字（ひらがな1文字など）',
  picture_card VARCHAR(255) NOT NULL COMMENT '絵札画像のURLまたはパス',
  text VARCHAR(255) NOT NULL COMMENT 'かるたの文章',
  text_card VARCHAR(255) NOT NULL COMMENT '読み札画像のURLまたはパス',
  yomi VARCHAR(255) NOT NULL COMMENT '読み仮名',
  score INT NOT NULL DEFAULT 1 COMMENT '札の得点',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  CONSTRAINT pk_cards PRIMARY KEY (id)
) COMMENT='かるたのカード情報';