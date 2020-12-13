package twelve

var (
	number = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth"}

	gift = []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, "}
)

// Song returns a song.
func Song() string {
	song := ""
	sep := ""
	for i, _ := range number {
		song = song + sep + Verse(i+1)
		sep = "\n"
	}
	return song
}

// Verse returns i-th line of the song.
func Verse(i int) string {
	line := "On the " + number[i-1] + " day of Christmas my true love gave to me: "
	for g := i - 1; g > 0; g-- {
		line = line + gift[g]
	}
	if i > 1 {
		line = line + "and "
	}
	return line + gift[0]
}
