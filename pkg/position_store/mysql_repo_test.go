package position_store

import (
	"testing"

	"github.com/moiot/gravity/pkg/config"
	"github.com/moiot/gravity/pkg/mysql_test"
	"github.com/stretchr/testify/require"
)

func TestMysqlPositionRepo_GetPut(t *testing.T) {
	r := require.New(t)

	dbCfg := mysql_test.SourceDBConfig()

	repo, err := NewMySQLRepo(dbCfg, "")
	r.NoError(err)

	// delete it first
	r.NoError(repo.Delete(t.Name()))

	_, exist, err := repo.Get(t.Name())
	r.NoError(err)

	r.False(exist)

	// put first value
	position := Position{
		Name:        t.Name(),
		Stage:       config.Stream,
		ValueString: "test",
	}
	r.NoError(repo.Put(t.Name(), position))

	p, exist, err := repo.Get(t.Name())
	r.NoError(err)
	r.True(exist)
	r.Equal("test", p.ValueString)
	r.Equal(config.Stream, p.Stage)

	// put another value
	position.ValueString = "test2"
	r.NoError(repo.Put(t.Name(), position))

	p2, exist, err := repo.Get(t.Name())
	r.NoError(err)
	r.True(exist)
	r.Equal(p2.ValueString, "test2")

	// put an invalid value
	position.ValueString = ""
	err = repo.Put(t.Name(), position)
	r.NotNil(err)
}
