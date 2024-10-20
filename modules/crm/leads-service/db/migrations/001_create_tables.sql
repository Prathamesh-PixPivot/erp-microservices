-- Creating the leads table
CREATE TABLE leads (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    status VARCHAR(50) NOT NULL,
    assigned_to INT, -- Assuming this will reference a user from the user-service
    organization_id INT, -- This will reference an organization from the organization-service
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Creating an index on leads to optimize searches by email
CREATE INDEX idx_leads_email ON leads(email);
