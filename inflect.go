package inflect

import(
    "regexp"
    "strings"
)

type Rule struct {
    patten *regexp.Regexp
    replacement string
}

var uncountables []string
var plurals []*Rule
var singulars []*Rule
var humans []*Rule
var acronyms map[string]bool
var acronymMatcher *regexp.Regexp

func init(){
    uncountables = make([]string,0)    
}

func Uncountables() []string {
    return uncountables
}

func AddPlural(patten, replacement string) {
    r := new(Rule)
    r.patten = regexp.MustCompile(patten)
    r.replacement = replacement
    plurals = append(plurals, r)
}

func AddSingular(patten, replacement string) {
    r := new(Rule)
    r.patten = regexp.MustCompile(patten)
    r.replacement = replacement
    singulars = append(plurals, r)
}

func AddHuman(patten, replacement string) {
    r := new(Rule)
    r.patten = regexp.MustCompile(patten)
    r.replacement = replacement
    humans = append(plurals, r)
}

func AddIrregular(singular, plural string) {
    // remove existing singular/plural
    // if strings.ToUpper(plural[0:1]) == strings.ToUpper(singular[0:1]) {
    //     AddPlural()
    //     AddPlural()
    //     AddSingular()
    // } else {
    // 
    // }
}

func AddAcronym(word string) {
    // add acro to map
    acronyms[strings.ToLower(word)] = true
    // rebuild regex
    var match string
    for word,_ := range acronyms {
        match += "|" + word
    }
    acronymMatcher = regexp.MustCompile(match[1:])
}

func AddUncountable(word string) {
    uncountables = append(uncountables, word)
}

func Pluralize(word string) string {
    return ""
}

func Singularize(word string) string {
    return ""
}

func Capitalize(word string) string {
    return ""
}

func Camelize(word string) string {
    return ""
}

func CamelizeDownFirst(word string) string {
    return ""
}

func Titleize(word string) string {
    return ""
}

func Underscore(word string) string {
    return ""
}

func Humanize(word string) string {
    return ""
}

func ForeignKey(word string) string {
    return ""
}

func ForeignKeyCondensed(word string) string {
    return ""
}

func Tableize(word string) string {
    return ""
}

func Parameterize(word string) string {
    return ""
}

func ParameterizeJoin(word, sep string) string {
    return ""
}

func Typeify(word string) string {
    return ""
}

func Dasherize(word string) string {
    return ""
}

func Ordinalize(word string) string {
    return ""
}















