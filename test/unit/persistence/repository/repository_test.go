package repository_test

import (
	"errors"
	"github.com/hbttundar/diabuddy-api-infra/persistence/repository"
	infraerrors "github.com/hbttundar/diabuddy-errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type fakeResult struct {
	affected int64
	err      error
}

func (f fakeResult) RowsAffected() (int64, error) { return f.affected, f.err }
func (f fakeResult) LastInsertId() (int64, error) { return 0, nil }

func TestParseResult_Success(t *testing.T) {
	repo := repository.BaseRepository{}
	err := repo.ParseResult(fakeResult{affected: 1}, "update")
	assert.Nil(t, err)
}

func TestParseResult_NoRows(t *testing.T) {
	repo := repository.BaseRepository{}
	err := repo.ParseResult(fakeResult{affected: 0}, "delete")
	require.Error(t, err)
	assert.Equal(t, infraerrors.NotFoundErrorType, err.Type())
}

func TestParseResult_Err(t *testing.T) {
	repo := repository.BaseRepository{}
	err := repo.ParseResult(fakeResult{err: errors.New("boom")}, "delete")
	require.Error(t, err)
	assert.Equal(t, infraerrors.InternalServerErrorType, err.Type())
}
