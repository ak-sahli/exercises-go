package clock

var createClockCases = []struct {
	hours, minutes int
	expected       string
}{
	{8, 0, "08:00"},        // on the hour
	{11, 9, "11:09"},       // past the hour
	{24, 0, "00:00"},       // midnight is zero hours
	{25, 0, "01:00"},       // hour rolls over
	{100, 0, "04:00"},      // hour rolls over continuously
	{1, 60, "02:00"},       // sixty minutes is next hour
	{0, 160, "02:40"},      // minutes roll over
	{0, 1723, "04:43"},     // minutes roll over continuously
	{25, 160, "03:40"},     // hour and minutes roll over
	{201, 3001, "11:01"},   // hour and minutes roll over continuously
	{72, 8640, "00:00"},    // hour and minutes roll over to exactly midnight
	{-1, 15, "23:15"},      // negative hour
	{-25, 0, "23:00"},      // negative hour rolls over
	{-91, 0, "05:00"},      // negative hour rolls over continuously
	{1, -40, "00:20"},      // negative minutes
	{1, -160, "22:20"},     // negative minutes roll over
	{1, -4820, "16:40"},    // negative minutes roll over continuously
	{-25, -160, "20:20"},   // negative hour and minutes both roll over
	{-121, -5810, "22:10"}, // negative hour and minutes both roll over continuously
}

var addMinutesCase = []struct {
	hours, minutes, minutesToAdd int
	expected                     string
}{
	{10, 0, 3, "10:03"},     // add minutes
	{6, 41, 0, "06:41"},     // add no minutes
	{0, 45, 40, "01:25"},    // add to next hour
	{10, 0, 61, "11:01"},    // add more than one hour
	{0, 45, 160, "03:25"},   // add more than two hours with carry
	{23, 59, 2, "00:01"},    // add across midnight
	{5, 32, 1500, "06:32"},  // add more than one day (1500 min = 25 hrs)
	{1, 1, 3500, "11:21"},   // add more than two days
	{10, 3, -3, "10:00"},    // subtract minutes
	{10, 3, -30, "09:33"},   // subtract to previous hour
	{10, 3, -70, "08:53"},   // subtract more than an hour
	{0, 3, -4, "23:59"},     // subtract across midnight
	{0, 0, -160, "21:20"},   // subtract more than two hours
	{6, 15, -160, "03:35"},  // subtract more than two hours with borrow
	{5, 32, -1500, "04:32"}, // subtract more than one day (1500 min = 25 hrs)
	{2, 20, -3000, "00:20"}, // subtract more than two days
}
