package Command

import (
	"reflect"
	"strings"
)

//SeperateArgs will seperate strings
func SeperateArgs(argString string, validFlags ...interface{}) (map[byte]string, error) {
	argMap := make(map[byte]string)

	for _, flag := range validFlags {
		if reflect.TypeOf(flag).Kind() == reflect.String {
			flagString := flag.(string)
			argIndex := strings.Index(argString, "-"+flagString+" ")

			nextIndex := strings.Index(argString[argIndex+3:], "-")
			if nextIndex == -1 {
				argMap[argString[argIndex+1]] = argString[argIndex+3:]
			} else {
				argMap[argString[argIndex+1]] = strings.TrimSpace(argString[argIndex+3 : nextIndex+argIndex+3])
			}
		}
	}

	return argMap, nil
}
