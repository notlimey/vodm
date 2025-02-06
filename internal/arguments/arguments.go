package arguments

import (
	"fmt"
	"github.com/notlimey/vodm/internal/file"
)

type Arguments struct {
	Urls  []string
	Flags *Flags
}

func ParseArguments(args []string) Arguments {
	var urls []string
	flags := &Flags{
		Concurrent: false,
		Output:     "",
		Limit:      0,
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		nextArg := ""
		if i+1 < len(args) {
			nextArg = args[i+1]
		}

		isFile := file.ArgumentIsFile(arg)
		if isFile {
			fileUrls := file.GetUrlsFromFile(arg)
			urls = append(urls, fileUrls...)
			continue
		}

		if IsFlag(arg) {
			skipNext, err := ParseFlag(arg, nextArg, flags)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
				continue
			}
			if skipNext {
				i++ // Skip the next arg
			}
			continue
		}

		// If not a file or flag, treat as URL
		urls = append(urls, arg)
	}

	return Arguments{
		Urls:  urls,
		Flags: flags,
	}
}
