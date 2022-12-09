package graph

import "github.com/hugovallada/go-expert/graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
