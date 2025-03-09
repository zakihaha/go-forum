ALTER TABLE posts
    ADD CONSTRAINT fk_user_id_posts
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;