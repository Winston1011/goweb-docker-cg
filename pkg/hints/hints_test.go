package hints_test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"goweb-docker-cg/pkg/hints"
)

var DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
	DryRun: true,
})

type User struct {
	ID   uint
	Name string
}

func AssertSQL(t *testing.T, result *gorm.DB, sql string) {
	if result.Statement.SQL.String() != sql {
		t.Errorf("SQL expects: %v, got %v", sql, result.Statement.SQL.String())
	}
}

func TestHint(t *testing.T) {
	result := DB.Find(&User{})

	AssertSQL(t, result, "SELECT * FROM `users`")

	result = DB.Clauses(hints.New("hint")).Find(&User{})
	AssertSQL(t, result, "SELECT /*+ hint */ * FROM `users`")

	result = DB.Clauses(hints.NewHint("/*#mode=READWRITE*/")).Find(&User{})
	AssertSQL(t, result, "SELECT /*#mode=READWRITE*/ * FROM `users`")

	result = DB.Clauses(hints.Comment("select", "hint")).Find(&User{})
	AssertSQL(t, result, "SELECT /* hint */ * FROM `users`")

	result = DB.Clauses(hints.CommentBefore("select", "hint")).Find(&User{})
	AssertSQL(t, result, "/* hint */ SELECT * FROM `users`")

	result = DB.Clauses(hints.CommentAfter("select", "hint")).Find(&User{})
	AssertSQL(t, result, "SELECT * /* hint */ FROM `users`")

	result = DB.Clauses(hints.CommentAfter("where", "hint")).Find(&User{}, "id = ?", 1)
	AssertSQL(t, result, "SELECT * FROM `users` WHERE id = ? /* hint */")
}

func TestIndexHint(t *testing.T) {
	result := DB.Clauses(hints.UseIndex("user_name")).Find(&User{})

	AssertSQL(t, result, "SELECT * FROM `users` USE INDEX (`user_name`)")

	result = DB.Clauses(hints.ForceIndex("user_name", "user_id").ForJoin()).Find(&User{})

	AssertSQL(t, result, "SELECT * FROM `users` FORCE INDEX FOR JOIN (`user_name`,`user_id`)")

	result = DB.Clauses(
		hints.ForceIndex("user_name", "user_id").ForJoin(),
		hints.IgnoreIndex("user_name").ForGroupBy(),
	).Find(&User{})

	AssertSQL(t, result, "SELECT * FROM `users` FORCE INDEX FOR JOIN (`user_name`,`user_id`) IGNORE INDEX FOR GROUP BY (`user_name`)")
}
