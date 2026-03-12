ALTER TABLE user_profiles ADD CONSTRAINT unique_userId_profiles UNIQUE (user_id);
ALTER TABLE user_credentials ADD CONSTRAINT unique_userId_credentials UNIQUE (user_id);