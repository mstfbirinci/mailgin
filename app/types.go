package app

//Exception holds failure information when error is occured
type Exception struct {
	Error            string
	ExceptionMessage string
}

//Message holds email header information for sending mail
type Message struct {
	Recipients  []string `json:"recipients" binding:"required"`
	From        string   `json:"from" binding:"required"`
	Subject     string   `json:"subject" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	ContentType string   `json:"contentType" binding:"required"`
}

type (
	//Configuration holds config object
	Configuration struct {
		Smtp Smtp
		Seq  Seq
	}
	//Smtp holds smtp server information
	Smtp struct {
		SmtpServer string
		User       string
		Pass       string
		Port       int
		SslEnabled bool
	}
	//Seq holds information for logging
	Seq struct {
		Host   string
		Signal string
	}
)
