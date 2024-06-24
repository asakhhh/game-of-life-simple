package game

import (
	"flag"
	"fmt"
	"os"
)

func IsValidArg(s string) bool {
	if s == "--help" || s == "--verbose" || s == "--edges-portal" || s == "--fullscreen" || s == "--footprints" || s == "--colored" || s == "--alternative" {
		return true
	}
	if len(s) > 11 && s[:11] == "--delay-ms=" {
		return true
	}
	if len(s) > 7 && s[:7] == "--file=" {
		return true
	}
	if len(s) > 9 && s[:9] == "--random=" {
		return true
	}
	return false
}

func ParseArgs() {
	var t, errArgs []string
	t = append(t, os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		if IsValidArg(os.Args[i]) {
			t = append(t, os.Args[i])
		} else {
			errArgs = append(errArgs, os.Args[i])
		}
	}
	if len(errArgs) > 0 {
		fmt.Printf("The following arguments were incorrect: " + Red)
		for i, v := range errArgs {
			fmt.Printf("%s", v)
			if i != len(errArgs)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println("\n" + Reset)
	}

	os.Args = t

	FlagHelp = flag.Bool("help", false, "")
	FlagVerbose = flag.Bool("verbose", false, "")
	FlagDelayms = flag.Int("delay-ms", -1, "")
	FlagFile = flag.String("file", "", "")
	FlagEdgesPortal = flag.Bool("edges-portal", false, "")
	FlagRandom = flag.String("random", "", "")
	FlagFullscreen = flag.Bool("fullscreen", false, "")
	FlagFootprints = flag.Bool("footprints", false, "")
	FlagColored = flag.Bool("colored", false, "")
	FlagAlternative = flag.Bool("alternative", false, "")

	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 10 && os.Args[i][:11] == "--delay-ms=" {
			numstring := os.Args[i][11:]
			if len(numstring) == 0 || stringToInt(numstring) == -1 {
				os.Args[i] = os.Args[i][:11] + "-2"
			}
		}
	}

	flag.Parse()

	if *FlagHelp {
		if os.Args[1] == "--help" {
			return
		}
		*FlagHelp = false
	}

	if *FlagRandom != "" && *FlagFile != "" {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i][:7] == "--file=" {
				*FlagRandom = ""
				return
			} else if os.Args[i][:9] == "--random=" {
				*FlagFile = ""
				return
			}
		}
	}
}
