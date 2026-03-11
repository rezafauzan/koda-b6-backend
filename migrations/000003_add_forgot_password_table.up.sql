CREATE TABLE forgot_password (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    code_otp INT,
    expired_at TIMESTAMP DEFAULT NOW() + INTERVAL '5 minute',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)