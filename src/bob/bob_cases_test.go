package bob

var bobTestCases = []struct {
	phrase   string
	expected string
}{
	{
		"Tom-ay-to, tom-aaaah-to.",
		"Whatever.",
	},
	{
		"WATCH OUT!",
		"Whoa, chill out!",
	},
	{
		"FCECDFCAAB",
		"Whoa, chill out!",
	},
	{
		"Does this cryogenic chamber make me look fat?",
		"Sure.",
	},
	{
		"You are, what, like 15?",
		"Sure.",
	},
	{
		"fffbbcbeab?",
		"Sure.",
	},
	{
		"Let's go make out behind the gym!",
		"Whatever.",
	},
	{
		"It's OK if you don't want to go to the DMV.",
		"Whatever.",
	},
	{
		"WHAT THE HELL WERE YOU THINKING?",
		"Whoa, chill out!",
	},
	{
		"1, 2, 3 GO!",
		"Whoa, chill out!",
	},
	{
		"1, 2, 3",
		"Whatever.",
	},
	{
		"4?",
		"Sure.",
	},
	{
		"ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
		"Whoa, chill out!",
	},
	{
		"ÜMLÄÜTS!",
		"Whoa, chill out!",
	},
	{
		"ÜMLäÜTS!",
		"Whatever.",
	},
	{
		"I HATE YOU",
		"Whoa, chill out!",
	},
	{
		"Ending with ? means a question.",
		"Whatever.",
	},
	{
		":) ?",
		"Sure.",
	},
	{
		"Wait! Hang on. Are you going to be OK?",
		"Sure.",
	},
	{
		"",
		"Fine. Be that way!",
	},
	{
		"          ",
		"Fine. Be that way!",
	},
	{
		"\t\t\t\t\t\t\t\t\t\t",
		"Fine. Be that way!",
	},
	{
		"\nDoes this cryogenic chamber make me look fat?\nno",
		"Whatever.",
	},
	{
		"         hmmmmmmm...",
		"Whatever.",
	},
	{
		"Okay if like my  spacebar  quite a bit?   ",
		"Sure.",
	},
	{
		"\n\r \t\v\u00a0\u2002",
		"Fine. Be that way!",
	},
}
