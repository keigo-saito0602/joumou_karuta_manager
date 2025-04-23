INSERT INTO users (id, name, email, password, role)
VALUES (
  10,
  'テストユーザー',
  'test@example.com',
  '$2a$10$7N3jiXAYU0g7AI0BFq8nFu1wPM5TDppXH61/Blj0rH9c/J5JkzTIS', -- password123
  'user'
)
ON DUPLICATE KEY UPDATE name = VALUES(name);
