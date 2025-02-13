package _5_15

type Day int

const (
	_ Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func IsWeekend(day Day) bool {
	return day == Saturday || day == Sunday
}
