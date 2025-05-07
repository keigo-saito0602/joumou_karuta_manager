CREATE TABLE event_scores (
  id SERIAL COMMENT 'イベント用スコアID',
  name VARCHAR(255) NOT NULL COMMENT '参加者名',
  feeling TEXT COMMENT '感想',
  score INTEGER NOT NULL COMMENT 'スコア',
  cards_taken INTEGER NOT NULL COMMENT '取った札の枚数',
  fault_count INTEGER NOT NULL COMMENT 'お手つきの回数',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  CONSTRAINT pk_event_scores PRIMARY KEY (id)
) COMMENT='イベントでの参加者ごとのスコア記録';