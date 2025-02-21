package port_actor

import (
	"testing"

	"github.com/EvertonTomalok/ports-challenge/internal/ports"
	"github.com/EvertonTomalok/ports-challenge/internal/repositories"
	"github.com/stretchr/testify/assert"
)

const jsonFixturePath = "./fixtures/dummy_data.json"

func Test_JsonHandlerWithoutRepositoryLimiter(t *testing.T) {
	// no size limit is set, it will use the default one (very large number 2 ** 20)
	repo := repositories.NewMemDB()
	fakeSvc := ports.NewService(repo)
	fakeActor := NewJsonActor(fakeSvc)
	t.Run("saving all ports data", func(t *testing.T) {
		err := fakeActor.HandleUpsertStream(jsonFixturePath)
		assert.Nil(t, err)
	})

	t.Run("assert keys are present", func(t *testing.T) {
		mustExist := []string{"AEAJM", "AEAUH", "AEDXB"}
		for _, k := range mustExist {
			_, found := repo.Get(k)
			assert.True(t, found)
		}
	})
}

func Test_JsonHandlerWithRepositoryLimiter(t *testing.T) {
	// notice the max size for the database is set to 1, so it will
	// overwrite the default value. Only 1 item will be stored.
	repo := repositories.NewMemDB(repositories.WithMaxSize(1))
	fakeSvc := ports.NewService(repo)
	fakeActor := NewJsonActor(fakeSvc)

	t.Run("repository limited to a single port saved only", func(t *testing.T) {
		err := fakeActor.HandleUpsertStream(jsonFixturePath)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, repositories.MaxSizeAchievedErr)
	})

	t.Run("assert keys are present", func(t *testing.T) {
		mustExist := []string{"AEAJM"} // this is the only key added
		for _, k := range mustExist {
			_, found := repo.Get(k)
			assert.True(t, found)
		}
	})

	t.Run("assert keys are not present", func(t *testing.T) {
		notPresent := []string{"AEAUH", "AEDXB"} // these other two won't be added
		for _, k := range notPresent {
			_, found := repo.Get(k)
			assert.False(t, found)
		}
	})
}
