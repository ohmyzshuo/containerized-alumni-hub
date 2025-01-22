package errors

// custom errors
var (
    // Alumni-related errors
    EmailFormatIncorrectError      = 701  // Email format incorrect
    PhoneFormatIncorrectError      = 702  // Phone number format incorrect
    AlumniNotFoundError            = 703  // Alumni not found
    AlumniAlreadyExistsError       = 704  // Alumni already exists
    AlumniDataIncompleteError      = 705  // Alumni data incomplete

    // Publication-related errors
    PublicationNotFoundError       = 710  // Publication not found
    PublicationAlreadyExistsError  = 711  // Publication already exists
    PublicationDataIncompleteError = 712  // Publication data incomplete
    PublicationFormatIncorrectError= 713  // Publication format incorrect

    // Event-related errors
    EventNotFoundError             = 720  // Event not found
    EventAlreadyExistsError        = 721  // Event already exists
    EventDataIncompleteError       = 722  // Event data incomplete
    EventDateConflictError         = 723  // Event date conflict

    // Attendance related errors
    CheckInNotFoundError           = 730  // Check-in record not found
    CheckInAlreadyExistsError      = 731  // Check-in record already exists
    CheckInDataIncompleteError     = 732  // Check-in data incomplete
    CheckInUnauthorizedError       = 733  // Check-in unauthorized

    // General errors
    UnauthorizedError              = 740  // Unauthorized action
)
