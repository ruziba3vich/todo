DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name = 'tasks'
    ) THEN
        CREATE TABLE tasks (
            id SERIAL PRIMARY KEY,
            user_id UUID REFERENCES users(id),
            name VARCHAR(100),
            title TEXT,
            content TEXT,
            createdOn TIMESTAMP,
            completedOn TIMESTAMP,
            isCompleted BOOLEAN
        );
    END IF;
END$$;
