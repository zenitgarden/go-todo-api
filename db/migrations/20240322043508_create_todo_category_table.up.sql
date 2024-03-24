CREATE TABLE todo_category(
    todo_id VARCHAR(255) NOT NULL,
    category_id VARCHAR(255) NOT NULL,
    PRIMARY KEY(todo_id, category_id)
) ENGINE = InnoDB;