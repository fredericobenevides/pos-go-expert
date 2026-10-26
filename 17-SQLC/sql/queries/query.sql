-- name: ListCategories :many
SELECT * from categories;

-- name: GetCategoryByID :one
SELECT * from categories
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) 
VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price) 
VALUES (?, ?, ?, ?, ?);

-- name: ListCourse :many
SELECT c.*, ca.name as category_name 
FROM courses c 
    JOIN categories ca on c.category_id = ca.id;