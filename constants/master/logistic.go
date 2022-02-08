package master

const (
	LogisticSLATypeInstant = iota + 1
	LogisticSLATypeSameDay
	LogisticSLATypeNextDay
	LogisticSLATypePriority
	LogisticSLATypeRegular
)

const (
	LogisticSLATypeInstantNotes  string = "Estimasi Pengantaran 1 - 5 Jam"
	LogisticSLATypeSameDayNotes  string = "Estimasi Pengantaran 5 - 12 Jam"
	LogisticSLATypeNextDayNotes  string = "Estimasi Pengantaran 24 Jam"
	LogisticSLATypePriorityNotes string = "Estimasi Pengantaran 1 - 2 Hari"
	LogisticSLATypeRegularNotes  string = "Estimasi Pengantaran 3 - 5 Hari"
)

const (
	LogisticStatusNew = iota + 1
	LogisticStatusActive
	LogisticStatusHold
	LogisticStatusInactive
)

const (
	LogisticServiceStatusNew = iota + 1
	LogisticServiceStatusActive
	LogisticServiceStatusHold
	LogisticServiceStatusInactive
)
