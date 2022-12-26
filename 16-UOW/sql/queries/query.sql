-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: CreateCategoryAndReturn :execresult
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price) values (?, ?, ?, ?, ?);

-- name: ListCourses :many
SELECT c.*, ca.name as category FROM courses c join categories ca on c.category_id = ca.id;