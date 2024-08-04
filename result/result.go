package result

type CodeT uint

const (
	// api / request / function:
	Unknown         CodeT = 0
	Success         CodeT = 1
	Error           CodeT = 2
	Debug           CodeT = 3
	ValidationError CodeT = 4 // wrong input data, bad request
	Timeout         CodeT = 5
	Forbidden       CodeT = 6 // permission deny
	NeedRoot        CodeT = 7 // Root rights needed

	// resource / object:
	Nil           CodeT = 10 // object is nil, cant operate on Nil object
	NotFound      CodeT = 11
	CreateError   CodeT = 12
	DeleteError   CodeT = 13
	AlreadyExists CodeT = 14

	// login / session:
	// InvalidEmail        ResultT = 2
	// MissingTotpToken    ResultT = 3
	// InvalidTotpToken    ResultT = 4
	// InvalidLicenseKey   ResultT = 5
	// SessionIDInvalid  ResultT = 32
	// SessionIDNotFound ResultT = 33

	FistLogin CodeT = 30
	// UserIsBlocked ResultT = 31

	// InstalationNotFound ResultT = 40
	// InstalationBlocked  ResultT = 41

	// ActionNotFound       ResultT = 60
	// ClusterConfigIsEmpty ResultT = 61
)
