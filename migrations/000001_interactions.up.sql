-- Включаем расширение для генерации UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создаем перечисление для типов взаимодействий
CREATE TYPE interaction_type AS ENUM ('message_sent', 'phone_revealed');

-- Создаем таблицу пользователей
CREATE TABLE users (
  user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT,
  email TEXT,
  phone_num VARCHAR(15)
);

-- Создаем таблицу взаимодействий
CREATE TABLE interactions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(user_id),
  seller_id UUID REFERENCES users(user_id),
  ad_id UUID,
  interaction_type interaction_type,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


