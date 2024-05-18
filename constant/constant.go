package constant

import "os"

const ApplicationName = `Notes-cli`
const Version = `v4.0.0`

const DateFormat = "2006/01/02 15:04:05"

var HomeFolder, _ = os.UserHomeDir()
var FilePath = HomeFolder + Slash + FileName

const FileName = "notes-cli.json"
const FileRWPermissions = 0644

const TaskCompleted string = "✔"
const TaskIncompleted string = "✘"
const ItemNone string = "✉"

const ScopeRE = `(?:^|\s)(@\w[-\w.]*(?:\b|,))`
const TagRE = `(?:^|\s)(#\w[-\w.]*(?:\b|,))`
const IntegerRE = `\d+`
const AnyStringRE = `^.*$`

const Space = " "
const Empty = ""
const Pipe = "|"
const Slash = "/"

const HeaderType = "TYPE"
const HeaderContent = "CONTENT"
const HeaderScopes = "SCOPES"

const AddPrompt = "Add: "
const SearchPrompt = "search: "
