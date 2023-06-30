package uniconv

import (
	"os"
	"strings"
)

var replacements = map[string]string{
	"\u00e1\u00bb\u00a3": "ợ",
	"\u00e1\u00bb\u009b": "ớ",
	"\u00e1\u00bb\u0085": "ễ",
	"\u00e1\u00bb\u00ab": "ừ",
	"\u00e1\u00bb\u0097": "ỗ",
	"\u00e1\u00bb\u00b1": "ự",
	"\u00e1\u00bb\u0083": "ể",
	"\u00e1\u00bb\u00a7": "ủ",
	"\u00e1\u00ba\u00a7": "ầ",
	"\u00e1\u00bb\u008d": "ọ",
	"\u00e1\u00ba\u00a1": "ạ",
	"\u00e1\u00bb\u0093": "ồ",
	"\u00e1\u00ba\u00bf": "ế",
	"\u00e1\u00ba\u00a3": "ả",
	"\u00e1\u00bb\u0091": "ố",
	"\u00e1\u00ba\u00b7": "ặ",
	"\u00e1\u00ba\u00a5": "ấ",
	"\u00e1\u00ba\u00ad": "ậ",
	"\u00e1\u00bb\u00a9": "ứ",
	"\u00e1\u00bb\u00b7": "ỷ",
	"\u00e1\u00bb\u00b3": "ỳ",
	"\u00e1\u00bb\u0081": "ề",
	"\u00e1\u00bb\u008b": "ị",
	"\u00e1\u00bb\u0087": "ệ",
	"\u00e1\u00ba\u00af": "ắ",
	"\u00e1\u00bb\u00b9": "ỹ",
	"\u00e1\u00bb\u009d": "ờ",
	"\u00e2\u0080\u0093": "-",
	"\u00e2\u0080\u0099": "`",
	"\u00c3\u00ad":       "í",
	"\u00c3\u00a0":       "à",
	"\u00c6\u00a1":       "ơ",
	"\u00c5\u00a9":       "ũ",
	"\u00c6\u00b0":       "ư",
	"\u00c3\u00ba":       "ú",
	"\u00c3\u00a2":       "â",
	"\u00c3\u00aa":       "ê",
	"\u00c3\u00b9":       "ù",
	"\u00c3\u00a1":       "á",
	"\u00c3\u00b4":       "ô",
	"\u00c4\u0090":       "Đ",
	"\u00c4\u00a9":       "ĩ",
	"\u00c4\u0083":       "ă",
	"\u00c3\u00ac":       "ì",
	"\u00c3\u00b5":       "õ",
	"\u00c3\u00a3":       "ã",
	"\u00c3\u00a8":       "è",
	"\u00c3\u00bd":       "ý",
	"\u00c3\u0082":       "Â",
	"\u00c3\u0080":       "À",
	"\u00c3\u0081":       "Á",
	"\u00c3\u0083":       "Ã",
	"\u00c3\u0084":       "Ä",
	"\u00c3\u0085":       "Å",
	"\u00c3\u0086":       "Æ",
	"\u00c3\u00a4":       "ä",
	"\u00c3\u00a5":       "å",
	"\u00c3\u00a6":       "æ",
	"\u00c3\u0087":       "Ç",
	"\u00c3\u00a7":       "ç",
	"\u00c3\u0090":       "Ð",
	"\u00c3\u00b0":       "ð",
	"\u00c3\u0088":       "È",
	"\u00c3\u0089":       "É",
	"\u00c3\u008a":       "Ê",
	"\u00c3\u008b":       "Ë",
	"\u00c3\u00a9":       "é",
	"\u00c3\u00ab":       "ë",
	"\u00c3\u008c":       "Ì",
	"\u00c3\u008d":       "Í",
	"\u00c3\u008e":       "Î",
	"\u00c3\u008f":       "Ï",
	"\u00c3\u00ae":       "î",
	"\u00c3\u00af":       "ï",
	"\u00c3\u0091":       "Ñ",
	"\u00c3\u00b1":       "ñ",
	"\u00c3\u0092":       "Ò",
	"\u00c3\u0093":       "Ó",
	"\u00c3\u0094":       "Ô",
	"\u00c3\u0095":       "Õ",
	"\u00c3\u0096":       "Ö",
	"\u00c3\u0098":       "Ø",
	"\u00c5\u0092":       "Œ",
	"\u00c3\u00b2":       "ò",
	"\u00c3\u00b3":       "ó",
	"\u00c3\u00b6":       "ö",
	"\u00c3\u00b8":       "ø",
	"\u00c5\u0093":       "œ",
	"\u00c3\u0099":       "Ù",
	"\u00c3\u009a":       "Ú",
	"\u00c3\u009b":       "Û",
	"\u00c3\u009c":       "Ü",
	"\u00c3\u00bb":       "û",
	"\u00c3\u00bc":       "ü",
	"\u00c3\u009d":       "Ý",
	"\u00c5\u00b8":       "Ÿ",
	"\u00c3\u00bf":       "ÿ",
}

func ConvText(text string) string {
	for k, v := range replacements {
		text = strings.ReplaceAll(text, k, v)
	}

	return text
}

func ConvFileToFile(inputFile string, outputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	text := string(content)

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, []byte(ConvText(text)), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ConvFileToString(inputFile string, outputFile string) (string, error) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	text := string(content)

	return ConvText(text), nil
}
