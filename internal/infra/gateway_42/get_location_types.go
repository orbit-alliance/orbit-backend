package gateway_42

// LoggedDate representa uma data no formato "YYYY-MM-DD".
type LoggedDate = string

// LoggedDuration representa a quantidade de horas logadas no formato "hh:mm:ss.ssssss".
type LoggedDuration = string

// LocationResponse mapeia <dia → horas logadas>.
type LocationResponse = map[LoggedDate]LoggedDuration
