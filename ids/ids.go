package ids

import "github.com/oklog/ulid/v2"

type ULID struct {
	ulid.ULID
}

type (
	UserID ULID
	PostID ULID
)

type ULIDs interface {
	ULID | UserID | PostID
}

func MakeULID[T ULIDs]() T {
	return T{ulid.Make()}
}
