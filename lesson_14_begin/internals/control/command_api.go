package control

// CommandValidator validates command-specific arguments
// TODO: Create CommandValidator type, function, accepts json.RawMessage, returns error

// CommandProcessor processes command-specific arguments
// TODO: Create CommandProcessor type, function, accepts json.RawMessage, returns json.RawMessage + error

// Registry of valid commands with their validators and processors
// TODO create validCommands which is a map of string:struct{}
// TODO: Note the struct{} value has 2 fields - the func types from above!
// TODO: Define entry with key "shellcode:, and assign 2 functions to struct value
