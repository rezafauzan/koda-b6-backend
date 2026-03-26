CREATE TABLE IF NOT EXISTS forgot_password (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    code_otp INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)