package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hugovallada/go-expert/sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	Name: "Backend",
	// 	Description: sql.NullString{
	// 		String: "Backend Description",
	// 	},
	// 	ID: uuid.NewString(),
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(categories)

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:   "7d99b84d-5e03-4c5d-847e-103abd7c0806",
	// 	Name: "Backend Updated",
	// 	Description: sql.NullString{
	// 		String: "Backend desc updated",
	// 		Valid:  true,
	// 	},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	fmt.Println(category.Name, category.Description.String)
	// }

	err = queries.DeleteCategory(ctx, "7d99b84d-5e03-4c5d-847e-103abd7c0806")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, category.Description.String)
	}
}
