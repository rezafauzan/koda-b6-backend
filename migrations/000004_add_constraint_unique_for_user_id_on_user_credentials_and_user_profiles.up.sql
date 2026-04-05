ALTER TABLE user_profiles ADD CONSTRAINT unique_userId_profiles UNIQUE (user_id);
ALTER TABLE user_credentials ADD CONSTRAINT unique_userId_credentials UNIQUE (user_id);
ALTER TABLE user_credentials ADD CONSTRAINT unique_email_credentials UNIQUE (email);
ALTER TABLE user_credentials ADD CONSTRAINT unique_phone_credentials UNIQUE (phone);