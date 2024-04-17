DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name = 'users'
    ) THEN
        CREATE TABLE users (
            id UUID,
            username VARCHAR(64),
            password VARCHAR(64),
            PRIMARY KEY (id)
        );

    END IF;
END$$;
