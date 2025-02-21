package repositories

import (
	"encoding/json"
	"testing"

	"github.com/EvertonTomalok/ports-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
)

func parsePort(jsonRaw string) (domain.Port, error) {
	var portData domain.Port
	err := json.Unmarshal([]byte(jsonRaw), &portData)

	return portData, err
}

func Test_MemDB(t *testing.T) {
	key := "AEAJM"
	portRaw := `{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`
	portData, err := parsePort(portRaw)
	if err != nil {
		t.Fail()
		return
	}

	memDB := NewMemDB()
	t.Run("insert port", func(t *testing.T) {
		err := memDB.Upsert(key, portData)
		assert.Nil(t, err)

		value, found := memDB.Get(key)
		assert.True(t, found)
		assert.Equal(t, "Ajman", value.Name)
	})

	t.Run("update port", func(t *testing.T) {
		portData.Name = "Testing" // overwritten name

		err := memDB.Upsert(key, portData)
		assert.Nil(t, err)

		value, found := memDB.Get(key)
		assert.True(t, found)
		assert.Equal(t, "Testing", value.Name)
	})
}

func Test_MemDB_Limits(t *testing.T) {
	memDB := NewMemDB(WithMaxSize(1)) // set map to have max size 1

	type testCase struct {
		testName string
		portRaw  string
		key      string
		mustFail bool
	}
	testCases := []testCase{
		{
			testName: "first insert must work",
			portRaw:  `{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`,
			key:      "AEAJM",
			mustFail: false,
		},
		{
			testName: "second insert must fail",
			portRaw:  `{"name": "Abu Dhabi", "coordinates": [54.37, 24.47], "city": "Abu Dhabi", "province": "Abu Z¸aby [Abu Dhabi]", "country": "United Arab Emirates", "alias": [], "regions": [], "timezone": "Asia/Dubai", "unlocs": ["AEAUH"], "code": "52001"}`,
			key:      "AEAUH",
			mustFail: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			portData, err := parsePort(test.portRaw)
			if err != nil {
				t.Fail()
				return
			}
			err = memDB.Upsert(test.key, portData)
			if test.mustFail {
				assert.NotNil(t, err)
				assert.ErrorIs(t, err, MaxSizeAchievedErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
