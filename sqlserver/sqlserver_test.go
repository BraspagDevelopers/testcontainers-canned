package sqlserver

import (
	"context"
	"database/sql"
	"testing"

	canned "github.com/BraspagDevelopers/testcontainers-canned"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhenDoesNotProvideImage_ShouldReturnError(t *testing.T) {
	ctx := context.Background()

	c, err := CreateContainer(ctx, ContainerRequest{})
	assert.EqualError(t, err, "a password must be provided")
	assert.Nil(t, c)
}

func TestHappyScenario(t *testing.T) {
	ctx := context.Background()

	c, err := CreateContainer(ctx, ContainerRequest{
		Password: "Database@1234!",
		Logger:   canned.NewTestingLogger(t),
	})
	require.NoError(t, err)

	cs, err := c.GoConnectionString(ctx)
	require.NoError(t, err)

	db, err := sql.Open("sqlserver", cs)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE MyData ( ID int PRIMARY KEY IDENTITY(1,1) );")
	require.NoError(t, err)

	err = c.Shutdown(ctx)
	require.NoError(t, err)
}
