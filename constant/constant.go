package constant

import "os"

const ApplicationName = `Notes-cli`
const Version = `v4.0.0`

const DateFormat = "2006/01/02 15:04:05"

var HomeFolder, _ = os.UserHomeDir()
var FilePath = HomeFolder + StrSlash + FileName

const FileName = "notes-cli.json"
const FileRWPermissions = 0644

const TaskCompleted string = "✔"
const TaskIncompleted string = "✘"
const ItemNone string = "✉"

const ReScope = `(?:^|\s)(@\w[-\w.]*(?:\b|,))`
const ReTag = `(?:^|\s)(#\w[-\w.]*(?:\b|,))`
const ReInteger = `\d+`
const ReAnyString = `^.*$`

const StrSpace = " "
const StrEmpty = ""
const StrPipe = "|"
const StrSlash = "/"

const HeaderType = "TYPE"
const HeaderContent = "CONTENT"
const HeaderScopes = "SCOPES"

const HTitleSeparator = StrSpace + StrPipe + StrSpace
const HTitleTotalCount = "Total: "
const HTitleNoteCount = "N:"
const HTitleTaskCount = "T:"

const PromptAdd = "Add: "
const PromptSearch = "Search: "
