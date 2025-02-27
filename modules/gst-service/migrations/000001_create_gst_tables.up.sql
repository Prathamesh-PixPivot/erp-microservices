CREATE TABLE gstr1 (
    id SERIAL PRIMARY KEY,
    gstin VARCHAR(15) NOT NULL,
    return_period VARCHAR(10) NOT NULL,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    invoice_date DATE NOT NULL,
    taxable_value NUMERIC(15,2) NOT NULL,
    tax_amount NUMERIC(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gstr2a (
    id SERIAL PRIMARY KEY,
    gstin VARCHAR(15) NOT NULL,
    return_period VARCHAR(10) NOT NULL,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    invoice_date DATE NOT NULL,
    taxable_value NUMERIC(15,2) NOT NULL,
    tax_amount NUMERIC(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gstr3b (
    id SERIAL PRIMARY KEY,
    gstin VARCHAR(15) NOT NULL,
    return_period VARCHAR(10) NOT NULL,
    tax_liability NUMERIC(15,2) NOT NULL,
    itc_claimed NUMERIC(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gstr9 (
    id SERIAL PRIMARY KEY,
    gstin VARCHAR(15) NOT NULL,
    return_period VARCHAR(10) NOT NULL,
    total_turnover NUMERIC(20,2) NOT NULL,
    tax_paid NUMERIC(20,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gstr9c (
    id SERIAL PRIMARY KEY,
    gstin VARCHAR(15) NOT NULL,
    return_period VARCHAR(10) NOT NULL,
    audit_details TEXT NOT NULL,
    reconciliation_statement TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
