package formatText

import (
	"fmt"
	"strconv"
	"strings"
)

func GetTextInBoldFormat(text string) string {
	mapping := map[rune]rune{}
	normal := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bolded := "𝗔𝗕𝗖𝗗𝗘𝗙𝗚𝗛𝗜𝗝𝗞𝗟𝗠𝗡𝗢𝗣𝗤𝗥𝗦𝗧𝗨𝗩𝗪𝗫𝗬𝗭𝗮𝗯𝗰𝗱𝗲𝗳𝗴𝗵𝗶𝗷𝗸𝗹𝗺𝗻𝗼𝗽𝗾𝗿𝘀𝘁𝘂𝘃𝘄𝘅𝘆𝘇𝟬𝟭𝟮𝟯𝟰𝟱𝟲𝟳𝟴𝟵"

	for i, char := range normal {
		mapping[char] = rune(bolded[i])
	}

	var builder strings.Builder
	for _, char := range text {
		if boldChar, exists := mapping[char]; exists {
			builder.WriteRune(boldChar)
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}


func GetNumberInAmountFormat(n string) string {
	numFloat, err := strconv.ParseFloat(n, 64)
	if err != nil {
		return "Invalid number"
	}
	numInt := int(numFloat)

	number := fmt.Sprintf("%d", numInt)
	nLen := len(number)
	if nLen <= 3 {
		return number
	}
	var result strings.Builder
	pre := nLen % 3
	if pre > 0 {
		result.WriteString(number[:pre])
		if nLen > pre {
			result.WriteString(",")
		}
	}
	for i := pre; i < nLen; i += 3 {
		result.WriteString(number[i : i+3])
		if i+3 < nLen {
			result.WriteString(",")
		}
	}
	return result.String()
}

