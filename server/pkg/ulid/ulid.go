package ulid

import (
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"time"

	oklogulid "github.com/oklog/ulid/v2"
)

type ULID struct {
	id oklogulid.ULID
}

var entropy = oklogulid.Monotonic(rand.Reader, 0)

func New() ULID {
	id := oklogulid.MustNew(oklogulid.Timestamp(time.Now()), entropy)
	return ULID{id: id}
}

func Parse(s string) (ULID, error) {
	u, err := oklogulid.Parse(s)
	if err != nil {
		return ULID{}, err
	}
	return ULID{id: u}, nil
}

func (ulid ULID) String() string {
	return ulid.id.String()
}

var zeroULID oklogulid.ULID

func (ulid *ULID) IsZero() bool {
	if ulid == nil {
		return true
	}
	return ulid.id == zeroULID
}

func (ulid ULID) Value() (driver.Value, error) {
	if ulid.id == zeroULID {
		return nil, nil
	}
	return ulid.String(), nil
}

func (ulid *ULID) Scan(value any) error {
	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("failed to scan ULID: unexpected type %T", value)
	}

	u, err := oklogulid.Parse(str)
	if err != nil {
		return fmt.Errorf("invalid ULID format: %w", err)
	}

	*ulid = ULID{id: u}
	return nil
}
