CREATE TABLE clients (
    doc_type TEXT NOT NULL,
    doc_number TEXT NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    CONSTRAINT pk_clients PRIMARY KEY (doc_type, doc_number),
    CONSTRAINT uq_email UNIQUE (email)
);
