package server

// Syncraft's Application store, maintaining server state
type SyncraftFSM struct {
	KVStore map[string]string
}

// Initialize the application state after restart
func (s *SyncraftFSM) Initialize() {

}

// Take a snapshot of the Application state and persist
func (s *SyncraftFSM) TakeSnapshot() {

}

// Apply a command to the FSM state
func (s *SyncraftFSM) ApplyCommand(command *Command) {

}
