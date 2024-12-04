package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Language struct to hold the language data
type Language struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Note     string `json:"note,omitempty"`
}

type AlfredResult struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
}

type Response struct {
	Items []AlfredResult `json:"items"`
}

// Sample JSON data
var data = []byte(`[
    {"language": "Afar", "code": "aa"},
    {"language": "Abkhazian", "code": "ab"},
    {"language": "Avestan", "code": "ae"},
    {"language": "Afrikaans", "code": "af"},
    {"language": "Akan", "code": "ak"},
    {"language": "Amharic", "code": "am"},
    {"language": "Aragonese", "code": "an"},
    {"language": "Arabic", "code": "ar"},
    {"language": "Assamese", "code": "as"},
    {"language": "Avaric", "code": "av"},
    {"language": "Aymara", "code": "ay"},
    {"language": "Azerbaijani", "code": "az"},
    {"language": "Bashkir", "code": "ba"},
    {"language": "Belarusian", "code": "be"},
    {"language": "Bulgarian", "code": "bg"},
    {"language": "Bihari", "code": "bh"},
    {"language": "Bislama", "code": "bi"},
    {"language": "Bambara", "code": "bm"},
    {"language": "Bengali", "code": "bn"},
    {"language": "Tibetan", "code": "bo"},
    {"language": "Breton", "code": "br"},
    {"language": "Bosnian", "code": "bs"},
    {"language": "Catalan", "code": "ca"},
    {"language": "Chechen", "code": "ce"},
    {"language": "Chamorro", "code": "ch"},
    {"language": "Corsican", "code": "co"},
    {"language": "Cree", "code": "cr"},
    {"language": "Czech", "code": "cs"},
    {"language": "Church Slavic", "code": "cu"},
    {"language": "Chuvash", "code": "cv"},
    {"language": "Welsh", "code": "cy"},
    {"language": "Danish", "code": "da"},
    {"language": "German", "code": "de"},
    {"language": "Divehi", "code": "dv"},
    {"language": "Dzongkha", "code": "dz"},
    {"language": "Ewe", "code": "ee"},
    {"language": "Greek", "code": "el"},
    {"language": "English", "code": "en"},
    {"language": "Esperanto", "code": "eo"},
    {"language": "Spanish", "code": "es"},
    {"language": "Estonian", "code": "et"},
    {"language": "Basque", "code": "eu"},
    {"language": "Persian", "code": "fa"},
    {"language": "Fula", "code": "ff"},
    {"language": "Finnish", "code": "fi"},
    {"language": "Fijian", "code": "fj"},
    {"language": "Faroese", "code": "fo"},
    {"language": "French", "code": "fr"},
    {"language": "Western Frisian", "code": "fy"},
    {"language": "Irish", "code": "ga"},
    {"language": "Scottish Gaelic", "code": "gd"},
    {"language": "Galician", "code": "gl"},
    {"language": "Guarani", "code": "gn"},
    {"language": "Gujarati", "code": "gu"},
    {"language": "Manx", "code": "gv"},
    {"language": "Hausa", "code": "ha"},
    {"language": "Hebrew", "code": "he"},
    {"language": "Hindi", "code": "hi"},
    {"language": "Hiri Motu", "code": "ho"},
    {"language": "Croatian", "code": "hr"},
    {"language": "Haitian Creole", "code": "ht"},
    {"language": "Hungarian", "code": "hu"},
    {"language": "Armenian", "code": "hy"},
    {"language": "Herero", "code": "hz"},
    {"language": "Interlingua", "code": "ia"},
    {"language": "Indonesian", "code": "id"},
    {"language": "Interlingue", "code": "ie"},
    {"language": "Igbo", "code": "ig"},
    {"language": "Sichuan Yi", "code": "ii"},
    {"language": "Inupiaq", "code": "ik"},
    {"language": "Ido", "code": "io"},
    {"language": "Icelandic", "code": "is"},
    {"language": "Italian", "code": "it"},
    {"language": "Inuktitut", "code": "iu"},
    {"language": "Japanese", "code": "ja"},
    {"language": "Javanese", "code": "jv"},
    {"language": "Georgian", "code": "ka"},
    {"language": "Kongo", "code": "kg"},
    {"language": "Kikuyu", "code": "ki"},
    {"language": "Kuanyama", "code": "kj"},
    {"language": "Kazakh", "code": "kk"},
    {"language": "Greenlandic", "code": "kl"},
    {"language": "Khmer", "code": "km"},
    {"language": "Kannada", "code": "kn"},
    {"language": "Korean", "code": "ko"},
    {"language": "Kanuri", "code": "kr"},
    {"language": "Kashmiri", "code": "ks"},
    {"language": "Kurdish", "code": "ku"},
    {"language": "Komi", "code": "kv"},
    {"language": "Cornish", "code": "kw"},
    {"language": "Kyrgyz", "code": "ky"},
    {"language": "Latin", "code": "la"},
    {"language": "Luxembourgish", "code": "lb"},
    {"language": "Ganda", "code": "lg"},
    {"language": "Limburgish", "code": "li"},
    {"language": "Lingala", "code": "ln"},
    {"language": "Lao", "code": "lo"},
    {"language": "Lithuanian", "code": "lt"},
    {"language": "Luba-Katanga", "code": "lu"},
    {"language": "Latvian", "code": "lv"},
    {"language": "Malagasy", "code": "mg"},
    {"language": "Marshallese", "code": "mh"},
    {"language": "Maori", "code": "mi"},
    {"language": "Macedonian", "code": "mk"},
    {"language": "Malayalam", "code": "ml"},
    {"language": "Mongolian", "code": "mn"},
    {"language": "Moldavian", "code": "mo"},
    {"language": "Marathi", "code": "mr"},
    {"language": "Malay", "code": "ms"},
    {"language": "Maltese", "code": "mt"},
    {"language": "Burmese", "code": "my"},
    {"language": "Nauru", "code": "na"},
    {"language": "Norwegian Bokmål", "code": "nb"},
    {"language": "North Ndebele", "code": "nd"},
    {"language": "Nepali", "code": "ne"},
    {"language": "Ndonga", "code": "ng"},
    {"language": "Dutch", "code": "nl"},
    {"language": "Norwegian Nynorsk", "code": "nn"},
    {"language": "Norwegian", "code": "no"},
    {"language": "South Ndebele", "code": "nr"},
    {"language": "Navajo", "code": "nv"},
    {"language": "Chichewa", "code": "ny"},
    {"language": "Occitan", "code": "oc"},
    {"language": "Ojibwa", "code": "oj"},
    {"language": "Oromo", "code": "om"},
    {"language": "Oriya", "code": "or"},
    {"language": "Ossetian", "code": "os"},
    {"language": "Punjabi", "code": "pa"},
    {"language": "Polish", "code": "pl"},
    {"language": "Saint Pierre and Miquelon", "code": "pm"},
    {"language": "Pitcairn", "code": "pn"},
    {"language": "Pashto", "code": "ps"},
    {"language": "Portuguese", "code": "pt"},
    {"language": "Quechua", "code": "qu"},
    {"language": "Romansh", "code": "rm"},
    {"language": "Rundi", "code": "rn"},
    {"language": "Romanian", "code": "ro"},
    {"language": "Russian", "code": "ru"},
    {"language": "Kinyarwanda", "code": "rw"},
	{"language": "Sanskrit", "code": "sa"},
    {"language": "Sardinian", "code": "sc"},
    {"language": "Sindhi", "code": "sd"},
    {"language": "Northern Sami", "code": "se"},
    {"language": "Sango", "code": "sg"},
    {"language": "Sinhala", "code": "si"},
    {"language": "Slovak", "code": "sk"},
    {"language": "Slovenian", "code": "sl"},
    {"language": "Samoan", "code": "sm"},
    {"language": "Shona", "code": "sn"},
    {"language": "Somali", "code": "so"},
    {"language": "Albanian", "code": "sq"},
    {"language": "Serbian", "code": "sr"},
    {"language": "Swati", "code": "ss"},
    {"language": "Southern Sotho", "code": "st"},
    {"language": "Sundanese", "code": "su"},
    {"language": "Swedish", "code": "sv"},
    {"language": "Swahili", "code": "sw"},
    {"language": "Tamil", "code": "ta"},
    {"language": "Telugu", "code": "te"},
    {"language": "Tajik", "code": "tg"},
    {"language": "Thai", "code": "th"},
    {"language": "Tigrinya", "code": "ti"},
    {"language": "Turkmen", "code": "tk"},
    {"language": "Tagalog", "code": "tl"},
    {"language": "Tswana", "code": "tn"},
    {"language": "Tonga", "code": "to"},
    {"language": "Turkish", "code": "tr"},
    {"language": "Tsonga", "code": "ts"},
    {"language": "Tatar", "code": "tt"},
    {"language": "Twi", "code": "tw"},
    {"language": "Tahitian", "code": "ty"},
    {"language": "Uighur", "code": "ug"},
    {"language": "Ukrainian", "code": "uk"},
    {"language": "Urdu", "code": "ur"},
    {"language": "Uzbek", "code": "uz"},
    {"language": "Venda", "code": "ve"},
    {"language": "Vietnamese", "code": "vi"},
    {"language": "Volapük", "code": "vo"},
    {"language": "Walloon", "code": "wa"},
    {"language": "Wolof", "code": "wo"},
    {"language": "Xhosa", "code": "xh"},
    {"language": "Yiddish", "code": "yi"},
    {"language": "Yoruba", "code": "yo"},
    {"language": "Zhuang", "code": "za"},
    {"language": "Zulu", "code": "zu"},
    {"language": "Chinese (Simplified)", "code": "zh-Hans"},
    {"language": "Chinese (Traditional)", "code": "zh-Hant"}
]`)

func main() {
	// Unmarshal the JSON data into a slice of Language
	var languages []Language
	if err := json.Unmarshal(data, &languages); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Check if a search term is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a search term.")
		return
	}

	// Get the search term from command line arguments
	searchTerm := os.Args[1]

	// 构建 Alfred 结果
	var alfredResults []AlfredResult
	// Perform fuzzy search
	for _, lang := range languages {
		if strings.Contains(strings.ToLower(lang.Language), strings.ToLower(searchTerm)) {
			alfredResult := AlfredResult{
				Title:        lang.Code,
				Subtitle:     lang.Language,
				Arg:          lang.Code,
				Autocomplete: fmt.Sprintf("%s %s", lang.Language, lang.Code),
			}
			alfredResults = append(alfredResults, alfredResult)
		}
	}

	if len(alfredResults) == 0 {
		fmt.Println("Please provide a search term.")
		return
	}

	// 构建响应
	response := Response{Items: alfredResults}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 输出结果
	fmt.Println(string(jsonResponse))
}
