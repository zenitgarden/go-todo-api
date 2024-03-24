ALTER TABLE todo_category ADD CONSTRAINT todo_category_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE todo_category ADD CONSTRAINT todo_category_todo_id_fkey FOREIGN KEY (todo_id) REFERENCES todos(id) ON DELETE CASCADE ON UPDATE CASCADE;