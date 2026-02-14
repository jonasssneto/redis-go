package entity

type Item struct {
	Value    string
	ExpireAt *int64
}

func (i *Item) IsExpired(now int64) bool {
	if i.ExpireAt == nil {
		return false
	}

	return now >= *i.ExpireAt
}
